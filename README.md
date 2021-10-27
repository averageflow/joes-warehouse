# Joe's Warehouse Software

[![PkgGoDev](https://pkg.go.dev/badge/github.com/averageflow/joes-warehouse)](https://pkg.go.dev/github.com/averageflow/joes-warehouse)
[![Go Report Card](https://goreportcard.com/badge/github.com/averageflow/joes-warehouse)](https://goreportcard.com/report/github.com/averageflow/joes-warehouse)
[![License](https://img.shields.io/github/license/averageflow/joes-warehouse.svg)](https://github.com/averageflow/joes-warehouse/blob/master/LICENSE.md)

Joe's Warehouse Software is a Go application that has the purpose of managing products and articles in your warehouse.

![https://github.com/averageflow/joes-warehouse/raw/main/web/assets/favicon/android-icon-144x144.png](https://github.com/averageflow/joes-warehouse/raw/main/web/assets/favicon/android-icon-144x144.png)

## Table of contents

* [Functionalities](#functionalities)
* [Running the application](#running-the-application)
* [Running for development](#running-for-development)
* [Additional information](#additional-information)
* [Technologies used](#technologies-used)
* [Why Go ?](#why-go-)
* [Unit tests](#unit-tests)
* [Possible Improvements](#possible-improvements)
* [Credits](#credits)

## Functionalities
In summary the application can:
* Create new products (from JSON request or from file upload)
* Create new articles (from JSON request or from file upload)
* Retrieve list of products present in warehouse (with related articles)
* Retrieve list of articles present in warehouse
* Sell products if stock is enough, or product is unlimited stock like software (from JSON request or from form)
* Log every sale of products (transactions)
* Retrieve list of transactions
* Delete products from the warehouse
* Delete articles from the warehouse

## Running the application
To kickstart the application and all dependencies required for its operation, you should be running on a machine with Docker installed.
Clone the project, or download the zip file with the source code from [the releases page](https://github.com/averageflow/joes-warehouse/releases) page. 
Then, from the root of the project run from the terminal:

```sh
# -d option to run as daemon in background
docker-compose up -d
```

The application runs on port `7000`.

You can force the images to be re-built with:

```sh
docker-compose up --build -d
```

### Running for development
If you would like to actively develop the application then you can run the application manually from source by running `go run main.go` from the `/cmd/joes-warehouse` folder, but you should still run the database provided in the Docker image. Keep in mind the application expects environment variables to be present, and thus you should consider having a `.env` file with the correct values, e.g.:

```sh
export APPLICATION_MODE="release"
export DATABASE_CONNECTION="postgres://user:pass@localhost:5432/db"
export WEB_ASSET_LOCATION="../../web"
```

You can load these variables into your environment with `source .env`. 
VSCode users will find a pre-made run and debug configuration and thus can run and debug the project from the IDE.

### Additional information
This application provides several endpoints for "headless" usage (without frontend) and also provides a frontend to ease the use.
Thus if we want to create new products / articles via an HTTP request with JSON body we use the normal endpoint. 
If we want to create new products / articles via uploading a file to a web-form then we use the UI.

Products are composed of 0 or more articles. Products that are composed of articles can be sold only if they are in stock. Products that are not composed of any article can always be sold. This is in order to take into account that the product is of "infinite stock".

A list of transactions performed (sales) that have occurred can be obtained via the API and with the frontend.

You can view the API specification by using the open source API client [Insomnia](https://insomnia.rest/) and opening the file at `/storage/http/joes_warehouse_http_spec.yaml` and learn how to use the application endpoints.

In order to use the UI you can simply visit [http://localhost:7000](http://localhost:7000) in your browser.

Some data files are present in `/storage/payload-files` that can be directly uploaded using the web forms. Bear in mind if you want to add new products (products.json files), the articles which compose the product should obviously already be present in the database (using inventory.json files).

This application includes a graceful shutdown mechanics and so whenever you stop it, or it receives a stop signal, it will first wait for any HTTP request currently being processed to be finished and then gracefully shutdown. This makes it possible to deploy it without downtime and to ensure a better experience for users.

A simple pagination system was added to the GET calls and works by using URL parameters, e.g. `http://localhost:7000/api/products?limit=100&offset=0`. The default pagination limit if not specified is 100 items. The default pagination limit for the frontend is 500 items.

The code has been written in an attempt to achieve as clean code as possible, with dependency injection of key components and with simplicity in mind, with no global state.

### Technologies used
This project was built using:
* [Go programming language (1.17+)](https://golang.org/)
    * [Gin Gonic web framework](https://github.com/gin-gonic/gin)
    * [Gomponents declarative HTML components](https://github.com/maragudk/gomponents)
    * [PGX PostgreSQL driver](https://github.com/jackc/pgx)
* [PostgreSQL database](https://www.postgresql.org/)
* [Bulma CSS framework](https://bulma.io/)
* [Docker](https://www.docker.com/)

### Why Go ?
This application is the perfect use case for using the Go programming language:
* Connect in a seamless way to a database, nice facilities for writing queries and communicating to the database
* Write type-safe compilable code, catch errors before they occur at runtime
* Incredible refactoring capabilities due to awesome type-safety
* Code simplicity and readability is great in Go, approaching foreign codebases becomes easier
* Testing is very powerful and baked into the language
* Easy to deploy, single binary applications
* Author's choice (me) by default for any project, unless good reasons justify not using it
* Super fast applications
* Great programming tool support

## Unit tests
You can run the unit tests for this project if you have Go installed, by at the root of the project executing:

```sh
go test -shuffle=on ./...
```

The unit tests will also be run every time the Docker image is rebuilt.

## Possible Improvements
Some compromises were made during development to simplify certain aspects and make the project quicker to develop. Find some suggestions for improvements below. When better defined, these should be turned into GitHub issues to better keep track of the progress and create separate branches for the features.

* Frontend pagination was deemed as out of scope for this project, but is a great improvement to consider. Currently the frontend is capped at 500 items. Since the pagination is fully controlled by the client, by means of URL parameters, and the backend will respond to the wanted limits and offsets, this is relatively easy to implement.
* The files provided contain a data structure that is not ideal for the task at hand, and thus some workarounds had to be made in order to support them. This includes some choices to the database schema, as well as in the application's code. For example providing the article id on creation does not seem a correct choice. Ideally these should be auto-incremented if possible.
* The API could have been designed to use UUIDs instead of numeric IDs since this provides several advantages, specially when clustering. It seemed to complicate things greatly though because the provided files contained numeric IDs, and then we would need to write all sorts of lookup functions, so this was deemed as out of scope for the project. The addition of UUIDs would not be too difficult though and would prove useful on a large scale system.
* The docker compose file contains "secrets" which for a production-ready application is not great. Either the file should be encrypted in a certain fashion or the secrets should be obtained from a Vault (Hashicorp Vault comes to mind).
* Some more security should be added in the forms, some CSRF token mechanism would mitigate many vulnerabilities.
* Authorization Bearer token is for now hardcoded into the application. Ideally a mechanism to generate and securely store API bearer tokens would be implemented, and in the middleware we would check if the provided token is valid one (perhaps also even including a different permission set per token). The current implementation already somehow secures the application and showcases a HTTP handler middleware.
* Frontend authorization and authentication would really be important, you don't want anyone to be able to edit the warehouse's items. The suggestion would be to have Cookie based authentication for the web interface. This makes sense specially for the sale of items. This should also be added as a column (perhaps user_id) to the "transactions" database table in order to be able to view who performed a transaction.
* A SPA (single page application) seemed as a lot of overhead for this simple project. It should be considered if more complicated behavior and state were to be added to the UI. For the scale of this project SSR (server side rendering) seemed like the natural choice and simplified the development, without compromising functionality. This is also in many ways more secure and compatible across browsers, simple HTML and forms. This frontend should be improved and should be showing more data than it does now, for a more useful system.
* The frontend might benefit from adding the related article list on a per product basis. The information is present so it is a matter of deciding the best way to show it.
* The addition of PATCH endpoints to modify some resources would be useful, then we could for example rename products and articles.
* The addition of delete buttons in the UI to remove some resources would be useful.
* Distributed tracing would be a good addition to the application specially if it were to communicate with more services in its operations. Personal choice would be [Jaeger](https://www.jaegertracing.io/).
* Some structured logging on errors would be a good addition, also in combination to adding the logs into the spans for tracing.

### Credits
The icon used for the repository and for the favicon was made by [Flat Icons](https://www.flaticon.com/authors/flat-icons).