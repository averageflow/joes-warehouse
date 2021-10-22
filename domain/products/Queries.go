package products

const getProductsQuery = `
	SELECT
		id,
		name,
		price
	FROM
		products
	WHERE
		id = ?;
`

const addProductsQuery = `
	INSERT INTO
		products (name, price)
	VALUES
		(?, ?);
`

const modifyProductByIDQuery = `
	UPDATE
		products
	SET
		name = ?,
		price = ?
	WHERE
		id = ?;
`

const deleteProductByIDQuery = `
	DELETE FROM
		products
	WHERE
		id = ?;
`
