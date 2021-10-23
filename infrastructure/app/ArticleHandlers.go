package app

import (
	"net/http"
	"time"

	"github.com/averageflow/joes-warehouse/domain/warehouse"
	"github.com/averageflow/joes-warehouse/infrastructure"
	"github.com/averageflow/joes-warehouse/infrastructure/views"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) getArticlesHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addArticlesHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody infrastructure.RawArticleUploadRequest

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ApplicationServerResponse{
				Message:       infrastructure.GetMessageForHTTPStatus(http.StatusUnprocessableEntity),
				Error:         err.Error(),
				UnixTimestamp: time.Now().Unix(),
			})

			return
		}

		parsedArticles := warehouse.ConvertRawArticle(requestBody.Inventory)
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

func (s *ApplicationServer) addArticlesFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		_ = views.ArticleSubmissionView().Render(c.Writer)
	}
}

func (s *ApplicationServer) modifyArticleHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) deleteArticleHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}
