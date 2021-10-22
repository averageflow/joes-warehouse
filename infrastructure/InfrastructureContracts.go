package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationHTTPHandler interface {
	Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	ServeHTTP(http.ResponseWriter, *http.Request)
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
}
