package warehouse

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/jackc/pgx/v4"
)

// GetFullProductResponse will return a list of products in the warehouse.
func GetFullProductResponse(db infrastructure.ApplicationDatabase, limit, offset int64) (*products.ProductResponseData, error) {
	productData, err := getProducts(db, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(productData.Data) == 0 {
		return &products.ProductResponseData{Data: make(map[int64]products.WebProduct), Sort: []int64{}}, nil
	}

	result, err := prepareProductDataResponse(db, productData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFullProductsByID will return a list of product information for the requested product IDs.
func GetFullProductsByID(db infrastructure.ApplicationDatabase, wantedProductIDs []int64) (*products.ProductResponseData, error) {
	productData, err := getProductsByID(db, wantedProductIDs)
	if err != nil {
		return nil, err
	}

	result, err := prepareProductDataResponse(db, productData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// prepareProductDataResponse takes raw results from database rows of products,
// load the articles for each product and calculate the stock of each product and
// update the data set.
func prepareProductDataResponse(db infrastructure.ApplicationDatabase,
	productData *products.ProductResponseData) (*products.ProductResponseData, error) {
	productIDs := products.CollectProductIDs(productData.Data)

	relatedArticles, err := getArticlesForProducts(db, productIDs)
	if err != nil {
		return nil, err
	}

	for i := range productData.Data {
		wantedProduct := productData.Data[i]
		wantedProduct.Articles = relatedArticles[i]
		wantedProduct.AmountInStock = products.ProductAmountInStock(wantedProduct)
		wantedProduct.IsInfiniteStock = len(wantedProduct.Articles) == 0
		productData.Data[i] = wantedProduct
	}

	result := products.ProductResponseData{
		Data: productData.Data,
		Sort: productData.Sort,
	}

	return &result, nil
}

// getProducts will return a list of products in the warehouse.
func getProducts(db infrastructure.ApplicationDatabase, limit, offset int64) (*products.ProductResponseData, error) {
	ctx := context.Background()
	rows, err := db.Query(
		ctx,
		products.GetProductsQuery,
		limit,
		offset,
	)

	return handleGetProductRows(rows, err)
}

// getProductsByID will return a list of product information for the requested product IDs.
func getProductsByID(db infrastructure.ApplicationDatabase, productIDs []int64) (*products.ProductResponseData, error) {
	ctx := context.Background()
	rows, err := db.Query(
		ctx,
		fmt.Sprintf(
			products.GetProductsByIDQuery,
			infrastructure.IntSliceToCommaSeparatedString(productIDs),
		),
	)

	return handleGetProductRows(rows, err)
}

// handleGetProductRows is the common logic to handle scanning rows from the `products` table.
func handleGetProductRows(rows pgx.Rows, err error) (*products.ProductResponseData, error) {
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	var productData []products.WebProduct

	for rows.Next() {
		var product products.WebProduct

		var rawPrice float64

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&rawPrice,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		roundedPrice, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rawPrice), 64)
		product.Price = roundedPrice

		productData = append(productData, product)
	}

	resultingProducts := make(map[int64]products.WebProduct, len(productData))
	sortProductData := make([]int64, len(productData))

	for i := range productData {
		resultingProducts[productData[i].ID] = productData[i]
		sortProductData[i] = productData[i].ID
	}

	result := products.ProductResponseData{
		Data: resultingProducts,
		Sort: sortProductData,
	}

	return &result, nil
}

// AddProducts will create new records in the `products` table.
func AddProducts(db infrastructure.ApplicationDatabase, productData []products.RawProduct) error {
	ctx := context.Background()

	now := time.Now().Unix()

	articleMap := make(map[int][]articles.ArticleProductRelation)

	for i := range productData {
		tx, err := db.Begin(ctx)
		if err != nil {
			_ = tx.Rollback(ctx)
			return err
		}

		var productID int

		err = tx.QueryRow(
			ctx,
			products.AddProductsQuery,
			productData[i].Name,
			productData[i].Price,
			now,
			now,
		).Scan(&productID)
		if err != nil {
			_ = tx.Rollback(ctx)
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}

		articleMap[productID] = articles.ConvertRawArticleFromProductFile(productData[i].Articles)
	}

	for i := range articleMap {
		if err := AddArticleProductRelation(db, i, articleMap[i]); err != nil {
			return err
		}
	}

	return nil
}

// SellProducts will coordinate a product sale. This decreases stocks of the several articles
// and creates a transaction record to log the sale event.
func SellProducts(db infrastructure.ApplicationDatabase, wantedProducts map[int64]int64) error {
	for i := range wantedProducts {
		if wantedProducts[i] < 1 {
			return products.ErrSaleFailedDueToIncorrectAmount
		}
	}

	transactionID, err := CreateTransaction(db)
	if err != nil {
		return err
	}

	productData, err := GetFullProductsByID(db, products.CollectProductIDsForSell(wantedProducts))
	if err != nil {
		return err
	}

	for i := range productData.Data {
		product := productData.Data[i]
		notEnoughArticleStock := product.AmountInStock < wantedProducts[i]

		if !product.IsInfiniteStock && notEnoughArticleStock {
			return products.ErrSaleFailedDueToInsufficientStock
		}
	}

	newStockMap := make(map[int64]int64)

	for i := range productData.Data {
		for j := range productData.Data[i].Articles {
			productArticle := productData.Data[i].Articles[j]

			requiredArticleAmountForSale := productArticle.AmountOf * wantedProducts[i]
			newArticleStock := productArticle.Stock - requiredArticleAmountForSale
			newStockMap[j] = newArticleStock
		}
	}

	if err := UpdateArticlesStocks(db, newStockMap); err != nil {
		return err
	}

	return CreateTransactionProductRelation(db, transactionID, wantedProducts)
}

// DeleteProducts will delete products and relations in the warehouse by ID.
func DeleteProducts(db infrastructure.ApplicationDatabase, productIDs []int64) error {
	if len(productIDs) == 0 {
		return nil
	}

	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		fmt.Sprintf(
			products.DeleteProductArticlesQuery,
			infrastructure.IntSliceToCommaSeparatedString(productIDs),
		),
	); err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		fmt.Sprintf(
			products.DeleteTransactionProductsQuery,
			infrastructure.IntSliceToCommaSeparatedString(productIDs),
		),
	); err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		fmt.Sprintf(
			products.DeleteProductQuery,
			infrastructure.IntSliceToCommaSeparatedString(productIDs),
		),
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
