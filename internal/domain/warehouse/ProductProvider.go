package warehouse

import (
	"context"
	"fmt"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/jackc/pgx/v4"
)

// GetFullProductResponse will return a list of products in the warehouse.
func GetFullProductResponse(db infrastructure.ApplicationDatabase) (*products.ProductResponseData, error) {
	productData, err := GetProducts(db)
	if err != nil {
		return nil, err
	}

	if len(productData.Data) == 0 {
		return nil, products.ErrNoProductsEmptyWarehouse
	}

	result, err := prepareProductDataResponse(db, productData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFullProductsByID will return a list of product information for the requested product IDs.
func GetFullProductsByID(db infrastructure.ApplicationDatabase, wantedProductIDs []int64) (*products.ProductResponseData, error) {
	productData, err := GetProductsByID(db, wantedProductIDs)
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

	relatedArticles, err := GetArticlesForProduct(db, productIDs)
	if err != nil {
		return nil, err
	}

	for i := range relatedArticles {
		wantedProduct := productData.Data[i]
		wantedProduct.Articles = relatedArticles[i]
		wantedProduct.AmountInStock = products.ProductAmountInStock(wantedProduct)
		productData.Data[i] = wantedProduct
	}

	result := products.ProductResponseData{
		Data: productData.Data,
		Sort: productData.Sort,
	}

	return &result, nil
}

func GetProducts(db infrastructure.ApplicationDatabase) (*products.ProductResponseData, error) {
	ctx := context.Background()
	rows, err := db.Query(
		ctx,
		products.GetProductsQuery,
	)

	return handleGetProductRows(rows, err)
}

func GetProductsByID(db infrastructure.ApplicationDatabase, productIDs []int64) (*products.ProductResponseData, error) {
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

func handleGetProductRows(rows pgx.Rows, err error) (*products.ProductResponseData, error) {
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	var productData []products.WebProduct

	for rows.Next() {
		var product products.WebProduct

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

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
		if productData.Data[i].AmountInStock < wantedProducts[i] {
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
