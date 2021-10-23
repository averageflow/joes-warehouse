package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetProducts(db infrastructure.ApplicationDatabase) (map[string]infrastructure.WebProductModel, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, getProductsQuery)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	var products []infrastructure.WebProductModel

	for rows.Next() {
		var product infrastructure.WebProductModel

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

	result := make(map[string]infrastructure.WebProductModel, len(products))

	for i := range products {
		result[products[i].UniqueID] = products[i]
	}

	return result, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, products []infrastructure.RawProductModel) error {
	ctx := context.Background()

	now := time.Now().Unix()

	articleMap := make(map[int][]infrastructure.ArticleProductRelationModel)

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

func ModifyProduct(product infrastructure.ProductModel) error {
	return nil
}

func DeleteProduct(product infrastructure.ProductModel) error {
	return nil
}
