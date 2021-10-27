package products

const GetProductsQuery = `
	SELECT
		id,
		item_name,
		price,
		created_at,
		updated_at
	FROM
		products
	ORDER BY
		id, created_at 
	LIMIT $1 OFFSET $2;
`

const GetProductsByIDQuery = `
	SELECT
		id,
		item_name,
		price,
		created_at,
		updated_at
	FROM
		products
	WHERE
		id IN (%s)
	ORDER BY
		id, created_At;
`

const AddProductsQuery = `
	INSERT INTO
		products (item_name, price, created_at, updated_at)
	VALUES
		($1, $2, $3, $4) RETURNING id;
`

const AddTransactionQuery = `
	INSERT INTO
		transactions (created_at)
	VALUES
		($1) RETURNING id;
`

const AddTransactionProductRelationQuery = `
	INSERT INTO
		transaction_products (
			transaction_id,
			product_id,
			amount_of,
			created_at
		)
	VALUES
		($1, $2, $3, $4);
`

const DeleteProductQuery = `
	DELETE FROM products WHERE id IN (%s);
`

const DeleteProductArticlesQuery = `
	DELETE FROM product_articles WHERE product_id IN (%s);
`

const DeleteTransactionProductsQuery = `
	DELETE FROM transaction_products WHERE product_id IN (%s);
`
