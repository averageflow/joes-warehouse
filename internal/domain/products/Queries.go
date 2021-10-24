package products

const GetProductsQuery = `
	SELECT
		id,
		item_name,
		price,
		created_at,
		updated_at
	FROM
		products;
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
		id IN (%s);
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
