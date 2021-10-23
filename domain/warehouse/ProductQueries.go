package warehouse

const getProductsQuery = `
select id, unique_id, item_name, price, created_at, updated_at from products;
`

const addProductsQuery = `
	INSERT INTO
		products (item_name, price, created_at, updated_at)
	VALUES
		($1, $2, $3, $4)
	RETURNING id;
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
