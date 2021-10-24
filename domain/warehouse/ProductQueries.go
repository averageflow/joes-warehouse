package warehouse

const getProductsQuery = `
select id, item_name, price, created_at, updated_at from products;
`

const getProductsByIDQuery = `
select id, item_name, price, created_at, updated_at from products WHERE id IN (%s);
`

const addProductsQuery = `
	INSERT INTO
		products (item_name, price, created_at, updated_at)
	VALUES
		($1, $2, $3, $4)
	RETURNING id;
`

const createTransactionQuery = `
	INSERT INTO transactions (created_at) VALUES ($1) RETURNING id;
`

const createTransactionProductRelationQuery = `
	INSERT INTO transaction_products (transaction_id, product_id, amount_of, created_at)
	VALUES ($1, $2, $3, $4);
`

// const modifyProductByUUIDQuery = `
// 	UPDATE
// 		products
// 	SET
// 	    item_name = $1,
// 		price = $2,
// 		updated_at = $3
// 	WHERE
// 		unique_id = $4;
// `

// const deleteProductByUUIDQuery = `
// 	DELETE FROM
// 		products
// 	WHERE
// 		unique_id = $1;
// `
