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

func (s *ApplicationServer) sellProductsHandler() func(*gin.Context) {
	type sellProductsRequest struct {
		Data map[int64]int64 `json:"data"`
	}

	return func(c *gin.Context) {
		var requestBody sellProductsRequest

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.SellProducts(s.State.DB, requestBody.Data); err != nil {
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

// func (s *ApplicationServer) modifyProductHandler() func(*gin.Context) {
// 	return func(c *gin.Context) {}
// }

// func (s *ApplicationServer) deleteProductHandler() func(*gin.Context) {
// 	return func(c *gin.Context) {}
// }
