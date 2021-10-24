package app

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/averageflow/joes-warehouse/domain/articles"
	"github.com/averageflow/joes-warehouse/domain/products"
	"github.com/averageflow/joes-warehouse/domain/warehouse"
	"github.com/averageflow/joes-warehouse/infrastructure/views"
	"github.com/gin-gonic/gin"
)

func handleBadFormSubmission(c *gin.Context) {
	c.Status(http.StatusBadRequest)
	_ = views.ErrorUploadingView().Render(c.Writer)
}

func handleBadSaleSubmission(c *gin.Context) {
	c.Status(http.StatusBadRequest)
	_ = views.ErrorSellingView().Render(c.Writer)
}

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

func (s *ApplicationServer) addArticlesFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		formFileContents, err := getFormFileContents(c)
		if err != nil {
			handleBadFormSubmission(c)
			return
		}

		var requestData articles.RawArticleUploadRequest

		if err := json.Unmarshal(formFileContents, &requestData); err != nil {
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

func (s *ApplicationServer) addProductsFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		formFileContents, err := getFormFileContents(c)
		if err != nil {
			handleBadFormSubmission(c)
			return
		}

		var requestData products.RawProductUploadRequest

		if err := json.Unmarshal(formFileContents, &requestData); err != nil {
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

func (s *ApplicationServer) sellProductFormHandler() func(*gin.Context) {
	type sellProductFormRequest struct {
		Amount    int64 `form:"amount"`
		ProductID int64 `form:"productID"`
	}

	return func(c *gin.Context) {
		var requestBody sellProductFormRequest

		if err := c.Bind(&requestBody); err != nil {
			handleBadSaleSubmission(c)
			return
		}

		convertedData := map[int64]int64{requestBody.ProductID: requestBody.Amount}
		if err := warehouse.SellProducts(s.State.DB, convertedData); err != nil {
			handleBadSaleSubmission(c)
			return
		}

		c.Status(http.StatusOK)
		_ = views.SuccessSellingView().Render(c.Writer)
	}
}
