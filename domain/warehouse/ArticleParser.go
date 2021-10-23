package warehouse

import (
	"strconv"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

const (
	wantedConversionBase = 10
	wantedConversionBits = 64
)

func ConvertRawArticle(articles []infrastructure.RawArticle) []infrastructure.Article {
	result := make([]infrastructure.Article, len(articles))

	for i := range articles {
		id, _ := strconv.ParseInt(articles[i].ID, wantedConversionBase, wantedConversionBits)
		stock, _ := strconv.ParseInt(articles[i].Stock, wantedConversionBase, wantedConversionBits)

		result[i] = infrastructure.Article{
			ID:    id,
			Name:  articles[i].Name,
			Stock: stock,
		}
	}

	return result
}

func ConvertRawArticleFromProductFile(articles []infrastructure.RawArticleFromProductFile) []infrastructure.ArticleProductRelation {
	result := make([]infrastructure.ArticleProductRelation, len(articles))

	for i := range articles {
		id, _ := strconv.ParseInt(articles[i].ID, wantedConversionBase, wantedConversionBits)
		amountOf, _ := strconv.ParseInt(articles[i].Stock, wantedConversionBase, wantedConversionBits)

		result[i] = infrastructure.ArticleProductRelation{
			ID:       id,
			AmountOf: amountOf,
		}
	}

	return result
}
