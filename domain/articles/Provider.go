package articles

type ArticleModel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type ArticleFromFileModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
}

type ArticleFromProductFileModel struct {
	ID    string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"amount_of"`
}
