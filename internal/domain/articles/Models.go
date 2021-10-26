package articles

type WebArticle struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Stock     int64  `json:"stock"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ArticleResponseData struct {
	Data map[int64]WebArticle
	Sort []int64
}

type ArticlesOfProductMap map[int64]map[int64]ArticleOfProduct

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

type GetArticlesHandlerResponse struct {
	Data map[int64]WebArticle `json:"data"`
	Sort []int64              `json:"sort"`
}
