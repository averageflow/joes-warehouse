package articles

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func GetArticlesForProduct() ([]ArticleModel, error) {
	return nil, nil
}

func AddArticles(db infrastructure.ApplicationDatabase, articles []ArticleModel) error {
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

func AddArticlesFromProductPayload(articles []LegacyArticleFromProductFileModel) error {
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

func ConvertLegacyArticleToStandard(articles []LegacyArticleModel) []ArticleModel {
	return nil
}

func ConvertLegacyArticleFromProductFileToStandard(articles []LegacyArticleFromProductFileModel) []ArticleModel {
	return nil
}
