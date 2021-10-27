package app

import (
	"log"
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

// getArticlesHandler will return a list of articles in the warehouse in JSON.
func (s *ApplicationServer) getArticlesHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		paginationDetails := s.getPaginationDetails(c)

		articleData, err := warehouse.GetArticles(s.State.DB, paginationDetails.Limit, paginationDetails.Offset)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		c.JSON(http.StatusOK, articles.GetArticlesHandlerResponse{
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
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		parsedArticles := articles.ConvertRawArticle(requestBody.Inventory)
		if err := warehouse.AddArticles(s.State.DB, parsedArticles); err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusInternalServerError),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.AddArticleStocks(s.State.DB, parsedArticles); err != nil {
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

// deleteArticleHandler deletes articles from the warehouse by ID specified in the URL.
func (s *ApplicationServer) deleteArticleHandler() func(*gin.Context) {
	type deleteArticleHandlerURI struct {
		ID int64 `uri:"id"`
	}

	return func(c *gin.Context) {
		var request deleteArticleHandlerURI

		if err := c.BindUri(&request); err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		if err := warehouse.DeleteArticles(s.State.DB, []int64{request.ID}); err != nil {
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
