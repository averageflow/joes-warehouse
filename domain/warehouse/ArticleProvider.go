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

func AddLegacyArticles(db infrastructure.ApplicationDatabase, articles []infrastructure.LegacyArticleModel) error {
	converted := ConvertLegacyArticleToStandard(articles)
	return AddArticles(db, converted)
}

func AddArticleProductRelation(db infrastructure.ApplicationDatabase, items infrastructure.LegacyProductUploadRequest) error {
	return nil
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

func ConvertLegacyArticleToStandard(articles []infrastructure.LegacyArticleModel) []infrastructure.ArticleModel {
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

func ConvertLegacyArticleFromProductFileToStandard(articles []infrastructure.LegacyArticleFromProductFileModel) []infrastructure.ArticleModel {
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
