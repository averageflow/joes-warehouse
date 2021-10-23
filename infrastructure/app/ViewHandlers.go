package app

import (
	"github.com/averageflow/joes-warehouse/infrastructure/views"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) homeViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		_ = views.HomeView().Render(c.Writer)
	}
}

func (s *ApplicationServer) addProductsFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		_ = views.ProductSubmissionView().Render(c.Writer)
	}
}

func (s *ApplicationServer) addArticlesFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		_ = views.ArticleSubmissionView().Render(c.Writer)
	}
}
