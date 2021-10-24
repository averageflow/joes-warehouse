package articles

type WebArticle struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Stock     int64  `json:"stock"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Article struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
}

type ArticleOfProduct struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AmountOf  int64  `json:"amount_of"`
	Stock     int64  `json:"stock"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ArticleProductRelation struct {
	ID       int64
	AmountOf int64
}

type RawArticle struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
}

type RawArticleFromProductFile struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"amount_of"`
}

type RawArticleUploadRequest struct {
	Inventory []RawArticle `json:"inventory"`
}
