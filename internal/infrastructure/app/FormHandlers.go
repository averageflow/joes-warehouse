package app

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure/views"
	"github.com/gin-gonic/gin"
)

// handleBadFormSubmission will show an error view on bad upload forms.
func handleBadFormSubmission(c *gin.Context) {
	c.Status(http.StatusBadRequest)
	_ = views.ErrorUploadingView().Render(c.Writer)
}

// handleBadSaleSubmission will show an error view on bad sale forms.
func handleBadSaleSubmission(c *gin.Context) {
	c.Status(http.StatusBadRequest)
	_ = views.ErrorSellingView().Render(c.Writer)
}

// getFormFileContents will extract data from multipart upload file forms into []byte.
func getFormFileContents(c *gin.Context) ([]byte, error) {
	file, err := c.FormFile("fileData")
	if err != nil {
		return nil, err
	}

	fileData, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer fileData.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, fileData); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// addArticlesFromFileHandler will process the upload form of an articles file.
func (s *ApplicationServer) addArticlesFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		formFileContents, err := getFormFileContents(c)
		if err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		var requestData articles.RawArticleUploadRequest

		if err := json.Unmarshal(formFileContents, &requestData); err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		parsedArticles := articles.ConvertRawArticle(requestData.Inventory)

		if err := warehouse.AddArticles(s.State.DB, parsedArticles); err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		if err := warehouse.AddArticleStocks(s.State.DB, parsedArticles); err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		c.Status(http.StatusOK)
		_ = views.SuccessUploadingView().Render(c.Writer)
	}
}

// addProductsFromFileHandler will process the upload form of a product file.
func (s *ApplicationServer) addProductsFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		formFileContents, err := getFormFileContents(c)
		if err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		var requestData products.RawProductUploadRequest

		if err := json.Unmarshal(formFileContents, &requestData); err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		if err := warehouse.AddProducts(s.State.DB, requestData.Products); err != nil {
			log.Println(err.Error())
			handleBadFormSubmission(c)

			return
		}

		c.Status(http.StatusOK)
		_ = views.SuccessUploadingView().Render(c.Writer)
	}
}

// sellProductFormHandler will process the sell product form.
func (s *ApplicationServer) sellProductFormHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody products.SellProductFormRequest

		if err := c.Bind(&requestBody); err != nil {
			log.Println(err.Error())
			handleBadSaleSubmission(c)

			return
		}

		convertedData := map[int64]int64{requestBody.ProductID: requestBody.Amount}
		if err := warehouse.SellProducts(s.State.DB, convertedData); err != nil {
			log.Println(err.Error())
			handleBadSaleSubmission(c)

			return
		}

		c.Status(http.StatusOK)
		_ = views.SuccessSellingView().Render(c.Writer)
	}
}
