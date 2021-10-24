package warehouse

import (
	"context"
	"log"
	"time"

	"github.com/averageflow/joes-warehouse/domain/articles"
	"github.com/averageflow/joes-warehouse/infrastructure"
)

func AddArticleStocks(db infrastructure.ApplicationDatabase, articleData []articles.Article) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
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
			return err
		}
	}

	return tx.Commit(ctx)
}

func UpdateArticlesStocks(db infrastructure.ApplicationDatabase, newStockMap map[int64]int64) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	for i := range newStockMap {
		if _, err := tx.Exec(
			ctx,
			articles.UpdateArticleStockQuery,
			newStockMap[i],
			i,
		); err != nil {
			log.Printf("update article id %d with stock %d", i, newStockMap[i])
			return err
		}
	}

	return tx.Commit(ctx)
}
