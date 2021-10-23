package warehouse

const getArticlesForProductQuery = `
	SELECT
		articles.id,
		articles.item_name
	FROM
		articles
		INNER JOIN product_articles ON product_articles.product_id = product.id
	WHERE
		product_id = $1;
`

const addArticlesForProductQuery = `
	INSERT INTO product_articles (article_id, product_id, created_at, updated_at) VALUES ($1, $2, $3, $4);
`

const addArticlesWithIDQuery = `
	INSERT INTO
		articles (id, item_name, created_at, updated_at)
	VALUES
		($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT articles_pkey DO UPDATE SET 
		item_name = $2, 
		updated_at = $4 
		WHERE articles.id = $1;
`

const addArticleStocksQuery = `
	INSERT INTO
		article_stocks (article_id, stock, created_at, updated_at)
	VALUES
		($1, $2, $3, $4);
`

const updateArticleByUUIDQuery = `
	UPDATE
		articles
	SET
	item_name = $1,
		updated_at = $2
	WHERE
		unique_id = $3;
`

const deleteArticleByUUIDQuery = `
	DELETE FROM
		articles
	WHERE
		unique_id = $1;
`

const modifyArticleStocksByIDQuery = `
	UPDATE article_stocks SET stock = $1, updated_at = $2 WHERE article_id = $3;
`
