package articles

import (
	"strconv"
)

const (
	wantedConversionBase = 10
	wantedConversionBits = 64
)

// ConvertRawArticle will convert articles originating from a request into
// database ready articles.
func ConvertRawArticle(articleData []RawArticle) []Article {
	result := make([]Article, len(articleData))

	for i := range articleData {
		id, _ := strconv.ParseInt(articleData[i].ID, wantedConversionBase, wantedConversionBits)
		stock, _ := strconv.ParseInt(articleData[i].Stock, wantedConversionBase, wantedConversionBits)

		result[i] = Article{
			ID:    id,
			Name:  articleData[i].Name,
			Stock: stock,
		}
	}

	return result
}

// ConvertRawArticleFromProductFile will convert articles originating from a product data into
// database ready articles.
func ConvertRawArticleFromProductFile(articleData []RawArticleFromProductFile) []ArticleProductRelation {
	result := make([]ArticleProductRelation, len(articleData))

	for i := range articleData {
		id, _ := strconv.ParseInt(articleData[i].ID, wantedConversionBase, wantedConversionBits)
		amountOf, _ := strconv.ParseInt(articleData[i].Stock, wantedConversionBase, wantedConversionBits)

		result[i] = ArticleProductRelation{
			ID:       id,
			AmountOf: amountOf,
		}
	}

	return result
}
