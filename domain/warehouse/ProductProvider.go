package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetFullProductResponse(db infrastructure.ApplicationDatabase) (map[string]infrastructure.WebProduct, []string, error) {
	products, sortProducts, err := GetProducts(db)
	if err != nil {
		return nil, nil, err
	}

	productIDs := CollectProductIDs(products)

	idtoUniqueIdMap := CollectProductIDsToUniqueIDs(products)

	relatedArticles, err := GetArticlesForProduct(db, productIDs)
	if err != nil {
		return nil, nil, err
	}

	for i := range relatedArticles {
		wantedUUID := idtoUniqueIdMap[i]
		wantedProduct := products[wantedUUID]
		wantedProduct.Articles = relatedArticles[i]
		products[wantedUUID] = wantedProduct
	}

	return products, sortProducts, nil
}

func GetProducts(db infrastructure.ApplicationDatabase) (map[string]infrastructure.WebProduct, []string, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, getProductsQuery)
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
			&product.UniqueID,
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

	result := make(map[string]infrastructure.WebProduct, len(products))
	orderData := make([]string, len(products))

	for i := range products {
		result[products[i].UniqueID] = products[i]
		orderData[i] = products[i].UniqueID
	}

	return result, orderData, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, products []infrastructure.RawProduct) error {
	ctx := context.Background()

	now := time.Now().Unix()

	articleMap := make(map[int][]infrastructure.ArticleProductRelation)

	for i := range products {
		tx, err := db.Begin(ctx)
		if err != nil {
			return err
		}

		var productID int

		err = tx.QueryRow(
			ctx,
			addProductsQuery,
			products[i].Name,
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

		articleMap[productID] = ConvertRawArticleFromProductFile(products[i].Articles)
	}

	for i := range articleMap {
		if err := AddArticleProductRelation(db, i, articleMap[i]); err != nil {
			return err
		}
	}

	return nil
}

func ModifyProduct(product infrastructure.Product) error {
	return nil
}

func DeleteProduct(product infrastructure.Product) error {
	return nil
}
