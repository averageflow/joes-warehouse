USE DATABASE joes_warehouse;

CREATE TABLE "articles" (
    "id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE "article_stocks" (
    "id" INTEGER NOT NULL,
    "article_id" INTEGER NOT NULL,
    "stock" INTEGER NOT NULL,
    FOREIGN KEY("article_id") REFERENCES "articles"("id"),
    PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE "products" (
    "id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "price" REAL NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE "product_articles" (
    "id" INTEGER NOT NULL,
    "product_id" INTEGER NOT NULL,
    "article_id" INTEGER NOT NULL,
    FOREIGN KEY("product_id") REFERENCES "products"("id"),
    FOREIGN KEY("article_id") REFERENCES "articles"("id"),
    PRIMARY KEY("id" AUTOINCREMENT)
);