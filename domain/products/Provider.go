package products

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetProducts() ([]ProductModel, error) {
	return nil, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, products []ProductModel) error {
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

func ModifyProduct(product ProductModel) error {
	return nil
}

func DeleteProduct(product ProductModel) error {
	return nil
}

func ConvertLegacyProductToStandard(products []LegacyProductModel) []ProductModel {
	var result []ProductModel
	return result
}
