package app

import (
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/domain/products"
	"github.com/averageflow/joes-warehouse/infrastructure"
	"github.com/averageflow/joes-warehouse/infrastructure/views"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) getProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addProductsHandler() func(*gin.Context) {
	type addProductsRequest struct {
		Data []products.ProductModel `json:"data"`
	}

	return func(c *gin.Context) {
		var requestBody addProductsRequest

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := products.AddProducts(s.State.DB, requestBody.Data); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		c.JSON(http.StatusOK, ApplicationServerResponse{
			Message:       infrastructure.GetMessageForHTTPStatus(http.StatusOK),
			UnixTimestamp: time.Now().Unix(),
		})
	}
}

func (s *ApplicationServer) addProductsFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addProductsFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		_ = views.ProductSubmissionView().Render(c.Writer)
	}
}

func (s *ApplicationServer) modifyProductHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) deleteProductHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}
