# Joe's Warehouse Software

Joe's Warehouse Software is a Go application that has the purpose of managing products and articles in your warehouse.

The application can be run by opening a terminal and navigating into `cmd/warehouse-manager-backend` and running:
```
go run main.go
```

## Additional notes

Some compromises were made during development to simplify certain aspects and make the project more portable and quicker to develop, namely:
* The use of a SQLite database in place of a full-blown database server eased the development and prototyping but for a more serious production implementation, a RDBMS like PostgreSQL or MariaDB should be used.
* The configuration file contains "secrets" which for a production-ready application should at the very least not be encrypted "on rest". Either the file should be encrypted in a certain fashion or the secrets should be obtained from a Vault (Hashicorp Vault comes to mind).

