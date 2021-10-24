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

func GetFullProductResponse(db infrastructure.ApplicationDatabase) (map[int64]products.WebProduct, []int64, error) {
	productData, sortProducts, err := GetProducts(db)
	if err != nil {
		return nil, nil, err
	}

	if len(productData) == 0 {
		return nil, nil, nil
	}

	productIDs := products.CollectProductIDs(productData)

	relatedArticles, err := GetArticlesForProduct(db, productIDs)
	if err != nil {
		return nil, nil, err
	}

	for i := range relatedArticles {
		wantedProduct := productData[i]
		wantedProduct.Articles = relatedArticles[i]
		wantedProduct.AmountInStock = products.ProductAmountInStock(wantedProduct)
		productData[i] = wantedProduct
	}

	return productData, sortProducts, nil
}

func GetFullProductsByID(db infrastructure.ApplicationDatabase, wantedProductIDs []int64) (map[int64]products.WebProduct, []int64, error) {
	productData, sortProducts, err := GetProductsByID(db, wantedProductIDs)
	if err != nil {
		return nil, nil, err
	}

	productIDs := products.CollectProductIDs(productData)

	relatedArticles, err := GetArticlesForProduct(db, productIDs)
	if err != nil {
		return nil, nil, err
	}

	for i := range relatedArticles {
		wantedProduct := productData[i]
		wantedProduct.Articles = relatedArticles[i]
		wantedProduct.AmountInStock = products.ProductAmountInStock(wantedProduct)
		productData[i] = wantedProduct
	}

	return productData, sortProducts, nil
}

func GetProducts(db infrastructure.ApplicationDatabase) (map[int64]products.WebProduct, []int64, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, products.GetProductsQuery)
	return handleGetProductRows(rows, err)
}

func GetProductsByID(db infrastructure.ApplicationDatabase, productIDs []int64) (map[int64]products.WebProduct, []int64, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, fmt.Sprintf(products.GetProductsByIDQuery, infrastructure.IntSliceToCommaSeparatedString(productIDs)))
	return handleGetProductRows(rows, err)
}

func handleGetProductRows(rows pgx.Rows, err error) (map[int64]products.WebProduct, []int64, error) {
	if err != nil {
		return nil, nil, err
	}

	if rows.Err() != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var productData []products.WebProduct

	for rows.Next() {
		var product products.WebProduct

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, nil, err
		}

		productData = append(productData, product)
	}

	result := make(map[int64]products.WebProduct, len(productData))
	sortProductData := make([]int64, len(productData))

	for i := range productData {
		result[productData[i].ID] = productData[i]
		sortProductData[i] = productData[i].ID
	}

	return result, sortProductData, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, productData []products.RawProduct) error {
	ctx := context.Background()

	now := time.Now().Unix()

	articleMap := make(map[int][]articles.ArticleProductRelation)

	for i := range productData {
		tx, err := db.Begin(ctx)
		if err != nil {
			return err
		}

		var productID int

		err = tx.QueryRow(
			ctx,
			products.AddProductsQuery,
			productData[i].Name,
			0,
			now,
			now,
		).Scan(&productID)
		if err != nil {
			return err
		}

		err = tx.Commit(ctx)
		if err != nil {
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
	transactionID, err := CreateTransaction(db)
	if err != nil {
		return err
	}

	productData, _, err := GetFullProductsByID(db, products.CollectProductIDsForSell(wantedProducts))
	if err != nil {
		return err
	}

	for i := range productData {
		if productData[i].AmountInStock < wantedProducts[i] {
			return products.ErrSaleFailedDueToInsufficientStock
		}
	}

	newStockMap := make(map[int64]int64)

	for i := range productData {
		for j := range productData[i].Articles {
			requiredArticleAmountForSale := productData[i].Articles[j].AmountOf * wantedProducts[i]
			newArticleStock := productData[i].Articles[j].Stock - requiredArticleAmountForSale
			newStockMap[j] = newArticleStock
		}
	}

	if err := UpdateArticlesStocks(db, newStockMap); err != nil {
		return err
	}

	return CreateTransactionProductRelation(db, transactionID, wantedProducts)
}