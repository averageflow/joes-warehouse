package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetProducts() ([]infrastructure.ProductModel, error) {
	return nil, nil
}

func AddProducts(db infrastructure.ApplicationDatabase, products []infrastructure.RawProductModel) error {
	ctx := context.Background()

	now := time.Now().Unix()

	articleMap := make(map[int][]infrastructure.ArticleModel)

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
