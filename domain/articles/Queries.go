package articles

const getArticlesForProductQuery = `
	SELECT
		articles.id,
		articles.name
	FROM
		articles
		INNER JOIN product_articles ON product_articles.product_id = product.id
	WHERE
		product_id = $1;
`

const addArticlesQuery = `
	INSERT INTO
		articles (name)
	VALUES
		($1);
`

const addArticleStocksQuery = `
	INSERT INTO
		article_stocks (article_id, stock)
	VALUES
		($1, $2);
`

const updateArticleByIDQuery = `
	UPDATE
		articles
	SET
		name = $1
	WHERE
		id = $2;
`

const deleteArticleByIDQuery = `
	DELETE FROM
		articles
	WHERE
		id = $1;
`

const modifyArticleStocksByIDQuery = `
	UPDATE article_stocks SET stock = $1 WHERE article_id = $2;
`
