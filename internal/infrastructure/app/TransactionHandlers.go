package app

import (
	"log"
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

// getTransactionsHandler returns a list of transactions in the warehouse in JSON format.
func (s *ApplicationServer) getTransactionsHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		paginationDetails := s.getPaginationDetails(c)

		transactionData, err := warehouse.GetTransactions(s.State.DB, paginationDetails.Limit, paginationDetails.Offset)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		c.JSON(http.StatusOK, transactionData)
	}
}
