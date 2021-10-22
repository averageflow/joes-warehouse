package products

import "github.com/averageflow/joes-warehouse/domain/articles"

type ProductModel struct {
	ID       int                     `json:"id"`
	Name     string                  `json:"name"`
	Price    float64                 `json:"price"`
	Articles []articles.ArticleModel `json:"articles"`
}

type LegacyProductModel struct {
	Name     string                                       `json:"name"`
	Articles []articles.LegacyArticleFromProductFileModel `json:"contain_articles"`
}
