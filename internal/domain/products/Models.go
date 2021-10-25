package products

import (
	"errors"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
)

var ErrSaleFailedDueToInsufficientStock = errors.New("sale failed, did not have enough stock for wanted product")
var ErrSaleFailedDueToIncorrectAmount = errors.New("sale failed, incorrect amount of products to sell was requested")
var ErrNoProductsEmptyWarehouse = errors.New("no products found due to empty warehouse")

type Product struct {
	ID       int64              `json:"id"`
	Name     string             `json:"name"`
	Price    float64            `json:"price"`
	Articles []articles.Article `json:"articles"`
}

type ProductResponseData struct {
	Data map[int64]WebProduct
	Sort []int64
}

type WebProduct struct {
	ID            int64                               `json:"id"`
	Name          string                              `json:"name"`
	Price         float64                             `json:"price"`
	AmountInStock int64                               `json:"amount_in_stock"`
	Articles      map[int64]articles.ArticleOfProduct `json:"articles"`
	CreatedAt     int64                               `json:"created_at"`
	UpdatedAt     int64                               `json:"updated_at"`
}

type RawProduct struct {
	Name     string                               `json:"name"`
	Price    float32                              `json:"price"`
	Articles []articles.RawArticleFromProductFile `json:"contain_articles"`
}

type RawProductUploadRequest struct {
	Products []RawProduct `json:"products"`
}
