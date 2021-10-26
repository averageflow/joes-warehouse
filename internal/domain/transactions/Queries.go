package transactions

const GetTransactionsQuery = `
	SELECT
		transactions.id,
		products.id AS product_id,
		products.item_name,
		transaction_products.amount_of,
		transactions.created_at
	FROM
		transactions
		INNER JOIN transaction_products ON transaction_products.transaction_id = transactions.id
		INNER JOIN products ON products.id = transaction_products.product_id
	ORDER BY
		transactions.id, transactions.created_at
	LIMIT $1 OFFSET $2;
`
