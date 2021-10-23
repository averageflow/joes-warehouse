package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetArticlesForProduct() (map[string]infrastructure.ArticleModel, error) {
	return nil, nil
}

func GetArticles(db infrastructure.ApplicationDatabase) (map[string]infrastructure.WebArticleModel, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, getArticlesQuery)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []infrastructure.WebArticleModel

	for rows.Next() {
		var article infrastructure.WebArticleModel

		err := rows.Scan(
			&article.ID,
			&article.UniqueID,
			&article.Name,
			&article.Stock,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	result := make(map[string]infrastructure.WebArticleModel, len(articles))

	for i := range articles {
		result[articles[i].UniqueID] = articles[i]
	}

	return result, nil
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
