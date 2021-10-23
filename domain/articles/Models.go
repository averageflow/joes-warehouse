package articles

type ArticleModel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type LegacyArticleModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
}

type LegacyArticleFromProductFileModel struct {
	ID    string `json:"art_id"`
	Stock string `json:"amount_of"`
}

type LegacyArticleUploadRequest struct {
	Products []LegacyArticleModel `json:"inventory"`
}
