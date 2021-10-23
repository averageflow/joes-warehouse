package warehouse

import (
	"strconv"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

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
