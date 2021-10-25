package app

import (
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

// getArticlesHandler will return a list of articles in the warehouse in JSON.
func (s *ApplicationServer) getArticlesHandler() func(*gin.Context) {
	type getArticlesHandlerResponse struct {
		Data map[int64]articles.WebArticle `json:"data"`
		Sort []int64                       `json:"sort"`
	}

	return func(c *gin.Context) {
		articleData, err := warehouse.GetArticles(s.State.DB)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		c.JSON(http.StatusOK, getArticlesHandlerResponse{
			Data: articleData.Data,
			Sort: articleData.Sort,
		})
	}
}

// addArticlesHandler adds articles to the warehouse from a JSON request body.
func (s *ApplicationServer) addArticlesHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody articles.RawArticleUploadRequest

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		parsedArticles := articles.ConvertRawArticle(requestBody.Inventory)
		if err := warehouse.AddArticles(s.State.DB, parsedArticles); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.AddArticleStocks(s.State.DB, parsedArticles); err != nil {
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
