package app

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

// getProductsHandler returns a list of products in the warehouse in JSON format.
func (s *ApplicationServer) getProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		paginationDetails := s.getPaginationDetails(c)

		productData, err := warehouse.GetFullProductResponse(s.State.DB, paginationDetails.Limit, paginationDetails.Offset)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		c.JSON(http.StatusOK, products.GetProductsHandlerResponse{
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
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.AddProducts(s.State.DB, requestBody.Products); err != nil {
			log.Println(err.Error())
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
			log.Println(err.Error())
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
			log.Println(err.Error())
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

// deleteProductHandler deletes products from the warehouse by ID specified in the URL.
func (s *ApplicationServer) deleteProductHandler() func(*gin.Context) {
	type deleteProductHandlerURI struct {
		ID int64 `uri:"id"`
	}

	return func(c *gin.Context) {
		var request deleteProductHandlerURI

		if err := c.BindUri(&request); err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.DeleteProducts(s.State.DB, []int64{request.ID}); err != nil {
			log.Println(err.Error())
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
