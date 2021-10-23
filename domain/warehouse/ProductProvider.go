package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetProducts() ([]infrastructure.ProductModel, error) {
	return nil, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, products []infrastructure.ProductModel) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range products {
		if _, err := tx.Exec(
			ctx,
			addProductsQuery,
			products[i].Name,
			products[i].Price,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func AddLegacyProducts(db infrastructure.ApplicationDatabase, products []infrastructure.LegacyProductModel) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range products {
		var productID int

		err := tx.QueryRow(
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

		err = AddArticles(db, ConvertLegacyArticleFromProductFileToStandard(products[i].Articles))
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func ModifyProduct(product infrastructure.ProductModel) error {
	return nil
}

func DeleteProduct(product infrastructure.ProductModel) error {
	return nil
}

func ConvertLegacyProductToStandard(products []infrastructure.LegacyProductModel) []infrastructure.ProductModel {
	result := make([]infrastructure.ProductModel, len(products))

	for i := range products {
		result[i] = infrastructure.ProductModel{
			Name: products[i].Name,
		}
	}

	return result
}
