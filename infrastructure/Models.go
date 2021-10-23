package infrastructure

type ArticleModel struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
}

type ArticleProductRelationModel struct {
	ID       int64
	AmountOf int64
}

type RawArticleModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
}

type RawArticleFromProductFileModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"amount_of"`
}

type RawArticleUploadRequest struct {
	Inventory []RawArticleModel `json:"inventory"`
}

type ProductModel struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Price    float64        `json:"price"`
	Articles []ArticleModel `json:"articles"`
}

type WebProductModel struct {
	ID        int64          `json:"id"`
	UniqueID  string         `json:"unique_id"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	Articles  []ArticleModel `json:"articles"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

type RawProductModel struct {
	Name     string                           `json:"name"`
	Articles []RawArticleFromProductFileModel `json:"contain_articles"`
}

type RawProductUploadRequest struct {
	Products []RawProductModel `json:"products"`
}
