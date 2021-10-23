package infrastructure

type ArticleModel struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
}

type LegacyArticleModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
}

type LegacyArticleFromProductFileModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"amount_of"`
}

type LegacyArticleUploadRequest struct {
	Inventory []LegacyArticleModel `json:"inventory"`
}

type ProductModel struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Price    float64        `json:"price"`
	Articles []ArticleModel `json:"articles"`
}

type LegacyProductModel struct {
	Name     string                              `json:"name"`
	Articles []LegacyArticleFromProductFileModel `json:"contain_articles"`
}

type LegacyProductUploadRequest struct {
	Products []LegacyProductModel `json:"products"`
}
