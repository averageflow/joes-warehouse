package warehouse

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/averageflow/joes-warehouse/domain/articles"
	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetArticlesForProduct(db infrastructure.ApplicationDatabase, productIDs []int64) (map[int64]map[int64]articles.ArticleOfProduct, error) {
	ctx := context.Background()

	rows, err := db.Query(
		ctx,
		fmt.Sprintf(getArticlesForProductQuery, infrastructure.IntSliceToCommaSeparatedString(productIDs)),
	)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	articleMap := make(map[int64]map[int64]articles.ArticleOfProduct)

	for rows.Next() {
		var article articles.ArticleOfProduct

		var productID int64

		err := rows.Scan(
			&productID,
			&article.ID,
			&article.Name,
			&article.AmountOf,
			&article.Stock,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		_, ok := articleMap[productID]
		if !ok {
			articleMap[productID] = make(map[int64]articles.ArticleOfProduct)
		}

		articleMap[productID][article.ID] = article
	}

	return articleMap, nil
}

func GetArticles(db infrastructure.ApplicationDatabase) (map[int64]articles.WebArticle, []int64, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, getArticlesQuery)
	if err != nil {
		return nil, nil, err
	}

	if rows.Err() != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var articleData []articles.WebArticle

	for rows.Next() {
		var article articles.WebArticle

		err := rows.Scan(
			&article.ID,
			&article.Name,
			&article.Stock,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, nil, err
		}

		articleData = append(articleData, article)
	}

	result := make(map[int64]articles.WebArticle, len(articleData))
	sortArticleData := make([]int64, len(articleData))

	for i := range articleData {
		result[articleData[i].ID] = articleData[i]
		sortArticleData[i] = articleData[i].ID
	}

	return result, sortArticleData, nil
}

func AddArticles(db infrastructure.ApplicationDatabase, articles []articles.Article) error {
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

func AddArticleProductRelation(db infrastructure.ApplicationDatabase, productID int, articles []articles.ArticleProductRelation) error {
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

func AddArticleStocks(db infrastructure.ApplicationDatabase, articles []articles.Article) error {
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

// func ModifyArticles() error {
// 	return nil
// }

// func DeleteArticles() error {
// 	return nil
// }

func UpdateArticlesStocks(db infrastructure.ApplicationDatabase, newStockMap map[int64]int64) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	for i := range newStockMap {
		if _, err := tx.Exec(
			ctx,
			updateArticleStockQuery,
			newStockMap[i],
			i,
		); err != nil {
			log.Printf("update article id %d with stock %d", i, newStockMap[i])
			return err
		}
	}

	return tx.Commit(ctx)
}
