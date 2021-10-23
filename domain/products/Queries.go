package products

const getProductsQuery = `
	SELECT
		id,
		name,
		price
	FROM
		products;
`

const addProductsQuery = `
	INSERT INTO
		products (name, price, created_at, updated_at)
	VALUES
		($1, $2, $3, $4);
`

const modifyProductByUUIDQuery = `
	UPDATE
		products
	SET
		name = $1,
		price = $2,
		updated_at = $3
	WHERE
		unique_id = $4;
`

const deleteProductByUUIDQuery = `
	DELETE FROM
		products
	WHERE
		unique_id = $1;
`
