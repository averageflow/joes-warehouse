package warehouse

import (
	"context"
	"fmt"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

// getArticlesForProducts will return the associated articles per product, for the given product IDs.
func getArticlesForProducts(db infrastructure.ApplicationDatabase, productIDs []int64) (articles.ArticlesOfProductMap, error) {
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
		return nil, rows.Err()
	}

	defer rows.Close()

	articleMap := make(articles.ArticlesOfProductMap)

	for rows.Next() {
		var article articles.ArticleOfProduct

		var productID int64

		if err := rows.Scan(
			&productID,
			&article.ID,
			&article.Name,
			&article.AmountOf,
			&article.Stock,
			&article.CreatedAt,
			&article.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if _, ok := articleMap[productID]; !ok {
			articleMap[productID] = make(map[int64]articles.ArticleOfProduct)
		}

		articleMap[productID][article.ID] = article
	}

	return articleMap, nil
}

// GetArticles will return a list of products in the warehouse.
func GetArticles(db infrastructure.ApplicationDatabase, limit, offset int64) (*articles.ArticleResponseData, error) {
	ctx := context.Background()

	rows, err := db.Query(
		ctx,
		articles.GetArticlesQuery,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	var articleData []articles.WebArticle

	for rows.Next() {
		var article articles.WebArticle

		if err := rows.Scan(
			&article.ID,
			&article.Name,
			&article.Stock,
			&article.CreatedAt,
			&article.UpdatedAt,
		); err != nil {
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

// AddArticles will create new records in the `articles` database.
func AddArticles(db infrastructure.ApplicationDatabase, articleData []articles.Article) error {
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
			articles.AddArticlesWithIDQuery,
			articleData[i].ID,
			articleData[i].Name,
			now,
			now,
		); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

// AddArticleProductRelation will create new records in the `product_articles` table.
func AddArticleProductRelation(db infrastructure.ApplicationDatabase, productID int, articleData []articles.ArticleProductRelation) error {
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
			articles.AddArticlesForProductQuery,
			articleData[i].ID,
			productID,
			articleData[i].AmountOf,
			now,
			now,
		); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

// DeleteArticles will delete articles and relations in the warehouse by ID.
func DeleteArticles(db infrastructure.ApplicationDatabase, articleIDs []int64) error {
	if len(articleIDs) == 0 {
		return nil
	}

	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		fmt.Sprintf(
			articles.DeleteArticleStocksQuery,
			infrastructure.IntSliceToCommaSeparatedString(articleIDs),
		),
	); err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		fmt.Sprintf(
			articles.DeleteProductArticlesQuery,
			infrastructure.IntSliceToCommaSeparatedString(articleIDs),
		),
	); err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		fmt.Sprintf(
			articles.DeleteArticlesQuery,
			infrastructure.IntSliceToCommaSeparatedString(articleIDs),
		),
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
