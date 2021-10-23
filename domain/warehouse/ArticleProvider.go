package warehouse

import (
	"context"
	"strconv"
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
			addArticlesQuery,
			articles[i].Name,
			now,
			now,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func AddArticlesWithPreMadeID(db infrastructure.ApplicationDatabase, articles []infrastructure.ArticleModel) error {
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

func AddArticleProductRelation(db infrastructure.ApplicationDatabase, productID int, articles []infrastructure.RawArticleFromProductFileModel) error {
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

const (
	wantedConversionBase = 10
	wantedConversionBits = 64
)

func ConvertRawArticle(articles []infrastructure.RawArticleModel) []infrastructure.ArticleModel {
	result := make([]infrastructure.ArticleModel, len(articles))

	for i := range articles {
		id, _ := strconv.ParseInt(articles[i].ID, wantedConversionBase, wantedConversionBits)
		stock, _ := strconv.ParseInt(articles[i].Stock, wantedConversionBase, wantedConversionBits)

		result[i] = infrastructure.ArticleModel{
			ID:    id,
			Name:  articles[i].Name,
			Stock: stock,
		}
	}

	return result
}

func ConvertRawArticleFromProductFile(articles []infrastructure.RawArticleFromProductFileModel) []infrastructure.ArticleModel {
	result := make([]infrastructure.ArticleModel, len(articles))

	for i := range articles {
		id, _ := strconv.ParseInt(articles[i].ID, wantedConversionBase, wantedConversionBits)
		stock, _ := strconv.ParseInt(articles[i].Stock, wantedConversionBase, wantedConversionBits)

		result[i] = infrastructure.ArticleModel{
			ID:    id,
			Name:  articles[i].Name,
			Stock: stock,
		}
	}

	return result
}
