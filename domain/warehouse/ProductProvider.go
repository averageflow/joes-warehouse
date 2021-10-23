package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func CollectProductIDs(products map[string]infrastructure.WebProduct) []int64 {
	var result []int64

	for i := range products {
		result = append(result, products[i].ID)
	}
	return result
}

func CollectProductIDsToUniqueIDs(products map[string]infrastructure.WebProduct) map[int64]string {
	result := make(map[int64]string)

	for i := range products {
		result[products[i].ID] = products[i].UniqueID
	}

	return result
}

func GetProducts(db infrastructure.ApplicationDatabase) (map[string]infrastructure.WebProduct, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, getProductsQuery)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
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
			return nil, err
		}

		products = append(products, product)
	}

	result := make(map[string]infrastructure.WebProduct, len(products))

	for i := range products {
		result[products[i].UniqueID] = products[i]
	}

	return result, nil
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
