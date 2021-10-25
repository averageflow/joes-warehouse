package warehouse

import (
	"context"
	"fmt"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

func GetArticlesForProduct(db infrastructure.ApplicationDatabase, productIDs []int64) (articles.ArticlesOfProductMap, error) {
	ctx := context.Background()

	rows, err := db.Query(
		ctx,
		fmt.Sprintf(
			articles.GetArticlesForProductQuery,
			infrastructure.IntSliceToCommaSeparatedString(productIDs),
		),
	)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	articleMap := make(articles.ArticlesOfProductMap)

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

func GetArticles(db infrastructure.ApplicationDatabase) (*articles.ArticleResponseData, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, articles.GetArticlesQuery)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
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
			return nil, err
		}

		articleData = append(articleData, article)
	}

	resultingArticles := make(map[int64]articles.WebArticle, len(articleData))
	sortArticleData := make([]int64, len(articleData))

	for i := range articleData {
		resultingArticles[articleData[i].ID] = articleData[i]
		sortArticleData[i] = articleData[i].ID
	}

	result := articles.ArticleResponseData{
		Data: resultingArticles,
		Sort: sortArticleData,
	}

	return &result, nil
}

func AddArticles(db infrastructure.ApplicationDatabase, articleData []articles.Article) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range articleData {
		if _, err := tx.Exec(
			ctx,
			articles.AddArticlesWithIDQuery,
			articleData[i].ID,
			articleData[i].Name,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func AddArticleProductRelation(db infrastructure.ApplicationDatabase, productID int, articleData []articles.ArticleProductRelation) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	for i := range articleData {
		if _, err := tx.Exec(
			ctx,
			articles.AddArticlesForProductQuery,
			articleData[i].ID,
			productID,
			articleData[i].AmountOf,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
