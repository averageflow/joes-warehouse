package app

import (
	"net/http"

	"github.com/averageflow/joes-warehouse/domain/warehouse"
	"github.com/averageflow/joes-warehouse/infrastructure/views"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) productViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		products, sortProducts, err := warehouse.GetFullProductResponse(s.State.DB)
		if err != nil {
			panic(err.Error())
		}

		c.Status(http.StatusOK)
		_ = views.ProductView(products, sortProducts).Render(c.Writer)
	}
}

func (s *ApplicationServer) articleViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {

		articles, sortArticles, err := warehouse.GetArticles(s.State.DB)
		if err != nil {
			panic(err.Error())
		}

		c.Status(http.StatusOK)
		_ = views.ArticleView(articles, sortArticles).Render(c.Writer)
	}
}

func (s *ApplicationServer) addProductsFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		_ = views.ProductSubmissionView().Render(c.Writer)
	}
}

func (s *ApplicationServer) addArticlesFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		_ = views.ArticleSubmissionView().Render(c.Writer)
	}
}
