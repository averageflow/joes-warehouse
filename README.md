# Joe's Warehouse Software
Joe's Warehouse Software is a Go application that has the purpose of managing products and articles in your warehouse.

## Running the application
To kickstart the application and all dependencies required for its operation, you should be running on a machine with Docker installed, and from the root of the project run (use `-d` option to run as daemon in background):

```sh
docker-compose up -d
```

The application runs on port `7000`.

Before using the application please make sure to run

## Tech Stack
This project was built using:
* [Go programming language](https://golang.org/)
    * [Gin Gonic web framework](https://github.com/gin-gonic/gin)
    * [Gomponents declarative HTML components](https://github.com/maragudk/gomponents)
    * [PGX PostgreSQL driver](https://github.com/jackc/pgx)
* [PostgreSQL database](https://www.postgresql.org/)
* [Bulma CSS framework](https://bulma.io/)
* [Docker](https://www.docker.com/)


## Functionalities
This application provides several endpoints for "headless" usage (without frontend) and also provides a frontend to ease the use.
Thus if we want to create new products / articles via an HTTP request with JSON body we use the normal endpoint. 
If we want to create new products / articles via uploading a file to a web-form then we use the UI.

This application includes a graceful shutdown mechanics and so whenever you stop it, or it receives a stop signal, it will first wait for any HTTP request currently being processed to be finished and then gracefully shutdown. This makes it possible to deploy it without downtime and to ensure a better experience for users.

## Why Go ?
This application is the perfect use case for using the Go programming language:
* Flexible enough to support JSON and forms communication
* Connect in a seamless way to a database, nice facilities for writing queries and communicating to the database
* Write type-safe compilable code, catch errors before they occur at runtime
* Incredible refactoring capabilities due to awesome type-safety
* Code simplicity and readability is great in Go, approaching foreign codebases becomes easier
* Testing is very powerful and baked into the language
* Easy to deploy, single binary applications
* Author's choice (me) by default for any project, unless good reasons justify not using it
* Super fast applications
* Great programming tool support

## Possible Improvements
Some compromises were made during development to simplify certain aspects and make the project quicker to develop, namely:
* The files provided contain a data structure that is not ideal for the task at hand, and thus some workarounds had to be made in order to support them. This includes some choices to the database schema, as well as in the application's code. For example providing the article id on creation does not seem a correct choice. Ideally these should be auto-incremented if possible.
* The API could have been designed to use UUIDs instead of numeric IDs since this provides several advantages, specially when clustering. It seemed to complicate things greatly though because the provided files contained numeric IDs, and then we would need to write all sorts of lookup functions, so this was deemed as out of scope for the project. The addition of UUIDs would not be too difficult though and would prove useful on a large scale system.
* The docker compose file contains "secrets" which for a production-ready application is not great. Either the file should be encrypted in a certain fashion or the secrets should be obtained from a Vault (Hashicorp Vault comes to mind).
* Some more security should be added in the upload forms, some CSRF token in the form if server side rendered as it currently is and perhaps some honeypot field to avoid any sort of bot.
* Authorization and Authentication would really be important, you don't want anyone to be able to edit the warehouse's items. The suggestion would be to have both API tokens for headless usage and Cookie based authentication for the web interface. This makes sense specially for the sale of items. This should also be added as a column (perhaps user_id) to the "transactions" database table in order to be able to view who performed a transaction.
* A SPA (single page application) seemed as a lot of overhead for this simple project. It should be considered if more complicated behavior and state were to be added to the UI. For the scale of this project SSR (server side rendering) seemed like the natural choice and simplified the development, without compromising functionality. This is also in many ways more secure and compatible across browsers, simple HTML and forms.
