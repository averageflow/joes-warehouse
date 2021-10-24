package warehouse

import (
	"context"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

func CreateTransaction(db infrastructure.ApplicationDatabase) (int64, error) {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
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
		return 0, err
	}

	return transactionID, tx.Commit(ctx)
}

func CreateTransactionProductRelation(db infrastructure.ApplicationDatabase, transactionID int64, productData map[int64]int64) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
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
			return err
		}
	}

	return tx.Commit(ctx)
}
