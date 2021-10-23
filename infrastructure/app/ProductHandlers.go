package app

import (
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/domain/warehouse"
	"github.com/averageflow/joes-warehouse/infrastructure"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) getProductsHandler() func(*gin.Context) {
	type getProductsHandlerResponse struct {
		Data map[string]infrastructure.WebProduct `json:"data"`
	}

	return func(c *gin.Context) {
		productData, err := warehouse.GetFullProductResponse(s.State.DB)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		c.JSON(http.StatusOK, getProductsHandlerResponse{
			Data: productData,
		})
	}
}

func (s *ApplicationServer) addProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody infrastructure.RawProductUploadRequest

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.AddProducts(s.State.DB, requestBody.Products); err != nil {
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

// func (s *ApplicationServer) modifyProductHandler() func(*gin.Context) {
// 	return func(c *gin.Context) {}
// }

// func (s *ApplicationServer) deleteProductHandler() func(*gin.Context) {
// 	return func(c *gin.Context) {}
// }
