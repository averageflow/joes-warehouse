# Joe's Warehouse Software

Joe's Warehouse Software is a Go application that has the purpose of managing products and articles in your warehouse.

The application can be run by opening a terminal and navigating into `cmd/joes-warehouse` and running:
```
go run main.go
```

This will launch the application running on port `7000`.

This application includes a graceful shutdown mechanics and so whenever you stop it, or it receives a stop signal, it will first wait for any HTTP request currently being processed to be finished and then gracefully shutdown. This makes it possible to deploy it without downtime and to ensure a better experience for users.

This application provides several endpoints.

A compatibility layer has been added to the endpoints to create products and articles that permits uploading JSON files with a "legacy" data structure to fit the application's more type safe approach. 
Thus if we want to create new products / articles via an HTTP request with JSON body we use the normal endpoint. If we want to create new products / articles via uploading a file to a web-form then we use the "/file-submission" endpoints.

## Additional notes

Some compromises were made during development to simplify certain aspects and make the project more portable and quicker to develop, namely:
* The use of a SQLite database in place of a full-blown database server eased the development and prototyping but for a more serious production implementation, a RDBMS like PostgreSQL or MariaDB should be used.
* The configuration file contains "secrets" which for a production-ready application should at the very least not be encrypted "on rest". Either the file should be encrypted in a certain fashion or the secrets should be obtained from a Vault (Hashicorp Vault comes to mind).

