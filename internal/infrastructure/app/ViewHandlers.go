package app

import (
	"net/http"

	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure/views"
	"github.com/gin-gonic/gin"
)

// productViewHandler will render the view that shows a product list.
func (s *ApplicationServer) productViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		productData, err := warehouse.GetFullProductResponse(s.State.DB)
		if err != nil {
			panic(err.Error())
		}

		c.Status(http.StatusOK)
		_ = views.ProductView(productData).Render(c.Writer)
	}
}

// articleViewHandler will render the view that shows an article list.
func (s *ApplicationServer) articleViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		articleData, err := warehouse.GetArticles(s.State.DB)
		if err != nil {
			panic(err.Error())
		}

		c.Status(http.StatusOK)
		_ = views.ArticleView(articleData).Render(c.Writer)
	}
}

// addProductsFromFileViewHandler will render the view that shows the product file upload form.
func (s *ApplicationServer) addProductsFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		_ = views.ProductSubmissionView().Render(c.Writer)
	}
}

// addArticlesFromFileViewHandler will render the view that shows the article file upload form.
func (s *ApplicationServer) addArticlesFromFileViewHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		_ = views.ArticleSubmissionView().Render(c.Writer)
	}
}
