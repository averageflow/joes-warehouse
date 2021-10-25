package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

// AddArticleStocks will create a new record in the `article_stocks` table.
func AddArticleStocks(db infrastructure.ApplicationDatabase, articleData []articles.Article) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	now := time.Now().Unix()

	for i := range articleData {
		if _, err := tx.Exec(
			ctx,
			articles.AddArticleStocksQuery,
			articleData[i].ID,
			articleData[i].Stock,
			now,
			now,
		); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

// UpdateArticlesStocks will update records in the `article_stocks` table.
func UpdateArticlesStocks(db infrastructure.ApplicationDatabase, newStockMap map[int64]int64) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	for i := range newStockMap {
		if _, err := tx.Exec(
			ctx,
			articles.UpdateArticleStockQuery,
			newStockMap[i],
			i,
		); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}
