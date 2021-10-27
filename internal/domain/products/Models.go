package products

import (
	"errors"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
)

var ErrSaleFailedDueToInsufficientStock = errors.New("sale failed, did not have enough stock for wanted product")
var ErrSaleFailedDueToIncorrectAmount = errors.New("sale failed, incorrect amount of products to sell was requested")

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
	ID              int64                               `json:"id"`
	Name            string                              `json:"name"`
	Price           float64                             `json:"price"`
	AmountInStock   int64                               `json:"amount_in_stock"`
	IsInfiniteStock bool                                `json:"is_infinite_stock"`
	Articles        map[int64]articles.ArticleOfProduct `json:"articles"`
	CreatedAt       int64                               `json:"created_at"`
	UpdatedAt       int64                               `json:"updated_at"`
}

type RawProduct struct {
	Name     string                               `json:"name"`
	Price    float32                              `json:"price"`
	Articles []articles.RawArticleFromProductFile `json:"contain_articles"`
}

type RawProductUploadRequest struct {
	Products []RawProduct `json:"products"`
}

type SellProductRequest struct {
	Data []SellProductRequestItems `json:"data"`
}

type SellProductRequestItems struct {
	ProductID int64 `json:"productID"`
	Amount    int64 `json:"amount"`
}

type SellProductFormRequest struct {
	ProductID int64 `form:"productID"`
	Amount    int64 `form:"amount"`
}

type GetProductsHandlerResponse struct {
	Data map[int64]WebProduct `json:"data"`
	Sort []int64              `json:"sort"`
}
