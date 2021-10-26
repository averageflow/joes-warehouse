package transactions

type TransactionDetails struct {
	ID            int64  `json:"id"`
	ProductID     int64  `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductAmount int64  `json:"product_amount"`
	CreatedAt     int64  `json:"created_at"`
}

type TransactionResponse struct {
	Data map[int64][]TransactionDetails `json:"data"`
	Sort []int64                        `json:"sort"`
}
