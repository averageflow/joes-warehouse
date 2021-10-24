package warehouse

const getArticlesForProductQuery = `
	SELECT
		product_articles.product_id,
		articles.id,
		articles.item_name,
		product_articles.amount_of,
		article_stocks.stock,
		articles.created_at,
		articles.updated_at
	FROM
		articles
		INNER JOIN product_articles ON product_articles.article_id = articles.id
		INNER JOIN article_stocks ON article_stocks.article_id = articles.id
	WHERE
		product_articles.product_id IN (%s);
`

const getArticlesQuery = `
select articles.id, item_name, article_stocks.stock , articles.created_at, articles.updated_at from articles
inner join article_stocks on article_stocks.article_id  = articles.id;
`

const addArticlesForProductQuery = `
	INSERT INTO product_articles (article_id, product_id, amount_of, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5);
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
		($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT article_stocks_article_id_key DO UPDATE SET 
		stock = $2, updated_at = $4;
`

// const updateArticleByUUIDQuery = `
// 	UPDATE
// 		articles
// 	SET
// 	item_name = $1,
// 		updated_at = $2
// 	WHERE
// 		unique_id = $3;
// `

// const deleteArticleByUUIDQuery = `
// 	DELETE FROM
// 		articles
// 	WHERE
// 		unique_id = $1;
// `

// const modifyArticleStocksByIDQuery = `
// 	UPDATE article_stocks SET stock = $1, updated_at = $2 WHERE article_id = $3;
// `
