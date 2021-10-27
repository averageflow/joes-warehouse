package articles

const GetArticlesForProductQuery = `
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
		product_articles.product_id IN (%s) 
		ORDER BY articles.id, articles.created_at;
`

const GetArticlesQuery = `
	SELECT
		articles.id,
		item_name,
		article_stocks.stock,
		articles.created_at,
		articles.updated_at
	FROM
		articles
		INNER JOIN article_stocks ON article_stocks.article_id = articles.id
		ORDER BY articles.id, articles.created_at LIMIT $1 OFFSET $2;
`

const AddArticlesForProductQuery = `
	INSERT INTO
		product_articles (
			article_id,
			product_id,
			amount_of,
			created_at,
			updated_at
		)
	VALUES
		($1, $2, $3, $4, $5);
`

const AddArticlesWithIDQuery = `
	INSERT INTO
		articles (id, item_name, created_at, updated_at)
	VALUES
		($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT articles_pkey DO
	UPDATE
	SET
		item_name = $2,
		updated_at = $4
	WHERE
		articles.id = $1;
`

const AddArticleStocksQuery = `
	INSERT INTO
		article_stocks (article_id, stock, created_at, updated_at)
	VALUES
		($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT article_stocks_article_id_key DO
	UPDATE
	SET
		stock = $2,
		updated_at = $4;
`

const UpdateArticleStockQuery = `
	UPDATE
		article_stocks
	SET
		stock = $1
	WHERE
		article_id = $2;
`

const DeleteArticlesQuery = `
	DELETE FROM articles WHERE id IN (%s);
`

const DeleteProductArticlesQuery = `
	DELETE FROM product_articles WHERE article_id IN (%s);
`

const DeleteArticleStocksQuery = `
	DELETE FROM article_stocks WHERE article_id IN (%s);
`
