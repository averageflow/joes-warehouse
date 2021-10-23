package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetArticlesForProduct() ([]infrastructure.ArticleModel, error) {
	return nil, nil
}

func AddArticles(db infrastructure.ApplicationDatabase, articles []infrastructure.ArticleModel) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range articles {
		if _, err := tx.Exec(
			ctx,
			addArticlesWithIDQuery,
			articles[i].ID,
			articles[i].Name,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func AddArticleProductRelation(db infrastructure.ApplicationDatabase, productID int, articles []infrastructure.ArticleProductRelationModel) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range articles {
		if _, err := tx.Exec(
			ctx,
			addArticlesForProductQuery,
			articles[i].ID,
			productID,
			articles[i].AmountOf,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func AddArticleStocks(db infrastructure.ApplicationDatabase, articles []infrastructure.ArticleModel) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range articles {
		if _, err := tx.Exec(
			ctx,
			addArticleStocksQuery,
			articles[i].ID,
			articles[i].Stock,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func ModifyArticles() error {
	return nil
}

func DeleteArticles() error {
	return nil
}

func ReserveArticleStock() error {
	return nil
}
