package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader           = "Authorization"
	dangerouslyHardcodedAuthToken = "387bf0c7-86dc-410e-ba05-1362cc1979ab-a6466675-fa39-46de-9e0b-d8b4bd94b52d"
)

// authTokenMiddleware is the middleware for the API calls that protects them with a Bearer token.
func (s *ApplicationServer) authTokenMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		authToken := c.GetHeader(authorizationHeader)
		if authToken != fmt.Sprintf("Bearer %s", dangerouslyHardcodedAuthToken) {
			// the request should be aborted with a 401 in case the token is not valid
			// for the moment, the token is hardcoded into the application
			c.AbortWithStatusJSON(http.StatusUnauthorized, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnauthorized),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		// if the auth token matched a valid one we will let the request further be processed
		// and be passed into the next handler
		c.Next()
	}
}
