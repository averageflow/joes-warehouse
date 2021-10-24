package warehouse

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/averageflow/joes-warehouse/domain/articles"
	"github.com/averageflow/joes-warehouse/domain/products"
	"github.com/averageflow/joes-warehouse/infrastructure"
	"github.com/jackc/pgx/v4"
)

func GetFullProductResponse(db infrastructure.ApplicationDatabase) (map[int64]infrastructure.WebProduct, []int64, error) {
	productData, sortProducts, err := GetProducts(db)
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

func GetFullProductsByID(db infrastructure.ApplicationDatabase, wantedProductIDs []int64) (map[int64]infrastructure.WebProduct, []int64, error) {
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

func GetProducts(db infrastructure.ApplicationDatabase) (map[int64]infrastructure.WebProduct, []int64, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, getProductsQuery)
	return handleGetProductRows(rows, err)
}

func GetProductsByID(db infrastructure.ApplicationDatabase, productIDs []int64) (map[int64]infrastructure.WebProduct, []int64, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, fmt.Sprintf(getProductsByIDQuery, infrastructure.IntSliceToCommaSeparatedString(productIDs)))
	return handleGetProductRows(rows, err)
}

func handleGetProductRows(rows pgx.Rows, err error) (map[int64]infrastructure.WebProduct, []int64, error) {
	if err != nil {
		return nil, nil, err
	}

	if rows.Err() != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var products []infrastructure.WebProduct

	for rows.Next() {
		var product infrastructure.WebProduct

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

		products = append(products, product)
	}

	result := make(map[int64]infrastructure.WebProduct, len(products))
	orderData := make([]int64, len(products))

	for i := range products {
		result[products[i].ID] = products[i]
		orderData[i] = products[i].ID
	}

	return result, orderData, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, productData []infrastructure.RawProduct) error {
	ctx := context.Background()

	now := time.Now().Unix()

	articleMap := make(map[int][]infrastructure.ArticleProductRelation)

	for i := range productData {
		tx, err := db.Begin(ctx)
		if err != nil {
			return err
		}

		var productID int

		err = tx.QueryRow(
			ctx,
			addProductsQuery,
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

	products, _, err := GetFullProductsByID(db, products.CollectProductIDsForSell(wantedProducts))
	if err != nil {
		return err
	}

	for i := range products {
		if products[i].AmountInStock < wantedProducts[i] {
			return errors.New("did not have enough stock for wanted product")
		}
	}

	newStockMap := make(map[int64]int64)
	for i := range products {
		for j := range products[i].Articles {
			requiredArticleAmountForSale := products[i].Articles[j].AmountOf * wantedProducts[i]
			newArticleStock := products[i].Articles[j].Stock - requiredArticleAmountForSale
			newStockMap[j] = newArticleStock
		}
	}

	err = UpdateArticlesStocks(db, newStockMap)
	if err != nil {
		return err
	}

	return CreateTransactionProductRelation(db, transactionID, wantedProducts)
}

func CreateTransaction(db infrastructure.ApplicationDatabase) (int64, error) {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	now := time.Now().Unix()

	var transactionID int64

	err = tx.QueryRow(
		ctx,
		createTransactionQuery,
		now,
	).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	return transactionID, tx.Commit(ctx)
}

func CreateTransactionProductRelation(db infrastructure.ApplicationDatabase, transactionID int64, products map[int64]int64) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range products {
		if _, err := tx.Exec(
			ctx,
			createTransactionProductRelationQuery,
			transactionID,
			i,
			products[i],
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func ModifyProduct(product infrastructure.Product) error {
	return nil
}

func DeleteProduct(product infrastructure.Product) error {
	return nil
}
