package articles

const getArticlesForProductQuery = `
	SELECT
		articles.id,
		articles.name
	FROM
		articles
		INNER JOIN product_articles ON product_articles.product_id = product.id;
`

const addArticlesQuery = `
	INSERT INTO
		articles (name)
	VALUES
		(?);
`

const addArticleStocksQuery = `
	INSERT INTO
		article_stocks (article_id, stock)
	VALUES
		(?, ?);
`

const updateArticleQuery = `
	UPDATE
		articles
	SET
		name = ?
	WHERE
		id = ?;
`

const deleteArticleQuery = `
	DELETE FROM
		articles
	WHERE
		id = ?;
`

const modifyArticleStocksQuery = `
	UPDATE article_stocks SET stock = ? WHERE article_id = ?;
`
