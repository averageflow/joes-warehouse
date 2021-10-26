package infrastructure

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// ApplicationHTTPHandler represents the application's router contract, any object implementing
// the contract can be used as the application's HTTP router.
type ApplicationHTTPHandler interface {
	Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	ServeHTTP(http.ResponseWriter, *http.Request)
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
	Static(relativePath, root string) gin.IRoutes
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
}

// ApplicationDatabase represents the application's database contract, any object implementing
// the contract can be used as the application's database.
type ApplicationDatabase interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
