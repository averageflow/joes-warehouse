package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

// getProductsHandler returns a list of products in the warehouse in JSON format.
func (s *ApplicationServer) getProductsHandler() func(*gin.Context) {
	type getProductsHandlerResponse struct {
		Data map[int64]products.WebProduct `json:"data"`
		Sort []int64                       `json:"sort"`
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
			Data: productData.Data,
			Sort: productData.Sort,
		})
	}
}

// addProductsHandler adds products to the warehouse from a JSON request body.
func (s *ApplicationServer) addProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody products.RawProductUploadRequest

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

// sellProductsHandler performs a product sale from a JSON request body.
func (s *ApplicationServer) sellProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody products.SellProductRequest

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		itemsToSell := make(map[int64]int64)

		for i := range requestBody.Data {
			item := requestBody.Data[i]
			itemsToSell[item.ProductID] = item.Amount
		}

		if err := warehouse.SellProducts(s.State.DB, itemsToSell); err != nil {
			isUnprocessableEntityError := errors.Is(err, products.ErrSaleFailedDueToIncorrectAmount) ||
				errors.Is(err, products.ErrSaleFailedDueToInsufficientStock)

			if isUnprocessableEntityError {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
					Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
					Error:         err.Error(),
					UnixTimestamp: time.Now().Unix(),
				})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
					Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
					Error:         err.Error(),
					UnixTimestamp: time.Now().Unix(),
				})
			}

			return
		}

		c.JSON(http.StatusOK, ApplicationServerResponse{
			Message:       infrastructure.GetMessageForHTTPStatus(http.StatusOK),
			UnixTimestamp: time.Now().Unix(),
		})
	}
}
