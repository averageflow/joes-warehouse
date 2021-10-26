package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/domain/transactions"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

// CreateTransaction will create a new record in the `transactions` table and return its ID.
func CreateTransaction(db infrastructure.ApplicationDatabase) (int64, error) {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return 0, err
	}

	now := time.Now().Unix()

	var transactionID int64

	err = tx.QueryRow(
		ctx,
		products.AddTransactionQuery,
		now,
	).Scan(&transactionID)
	if err != nil {
		_ = tx.Rollback(ctx)
		return 0, err
	}

	return transactionID, tx.Commit(ctx)
}

// CreateTransactionProductRelation will create a new record in the `transaction_products` table.
func CreateTransactionProductRelation(db infrastructure.ApplicationDatabase, transactionID int64, productData map[int64]int64) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	now := time.Now().Unix()

	for i := range productData {
		if _, err := tx.Exec(
			ctx,
			products.AddTransactionProductRelationQuery,
			transactionID,
			i,
			productData[i],
			now,
		); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

// GetTransactions will return a list of transactions in the warehouse.
func GetTransactions(db infrastructure.ApplicationDatabase, limit, offset int64) (*transactions.TransactionResponse, error) {
	ctx := context.Background()

	rows, err := db.Query(
		ctx,
		transactions.GetTransactionsQuery,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	var items []transactions.TransactionDetails

	for rows.Next() {
		var item transactions.TransactionDetails

		if err := rows.Scan(
			&item.ID,
			&item.ProductID,
			&item.ProductName,
			&item.ProductAmount,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	transactionData := make(map[int64][]transactions.TransactionDetails)

	sort := []int64{}

	for i := range items {
		sort = infrastructure.AppendIfMissingInt64(sort, items[i].ID)
		transactionData[items[i].ID] = append(transactionData[items[i].ID], items[i])
	}

	result := transactions.TransactionResponse{
		Data: transactionData,
		Sort: sort,
	}

	return &result, nil
}
