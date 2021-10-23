package products

const getProductsQuery = `
	SELECT
		id,
		name,
		price
	FROM
		products
	WHERE
		id = $1;
`

const addProductsQuery = `
	INSERT INTO
		products (name, price)
	VALUES
		($1, $2);
`

const modifyProductByIDQuery = `
	UPDATE
		products
	SET
		name = $1,
		price = $2
	WHERE
		id = $3;
`

const deleteProductByIDQuery = `
	DELETE FROM
		products
	WHERE
		id = $1;
`
