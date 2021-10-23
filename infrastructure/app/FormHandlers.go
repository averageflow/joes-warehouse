package app

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/averageflow/joes-warehouse/domain/warehouse"
	"github.com/averageflow/joes-warehouse/infrastructure"
	"github.com/averageflow/joes-warehouse/infrastructure/views"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) addDataFromFileHandler(itemType int) func(*gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile("fileData")
		if err != nil {
			handleBadFormSubmission(c)
			return
		}

		fileData, err := file.Open()
		if err != nil {
			handleBadFormSubmission(c)
			return
		}

		defer fileData.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, fileData); err != nil {
			handleBadFormSubmission(c)
			return
		}

		log.Println(buf.String())

		if itemType == infrastructure.ItemTypeArticle {
			var requestData infrastructure.LegacyArticleUploadRequest

			if err := json.Unmarshal(buf.Bytes(), &requestData); err != nil {
				handleBadFormSubmission(c)
				return
			}

			if err := warehouse.AddArticlesWithPreMadeID(s.State.DB, warehouse.ConvertLegacyArticleToStandard(requestData.Inventory)); err != nil {
				log.Println(err.Error())
				handleBadFormSubmission(c)
				return
			}

		} else if itemType == infrastructure.ItemTypeProduct {
			var requestData infrastructure.LegacyProductUploadRequest

			if err := json.Unmarshal(buf.Bytes(), &requestData); err != nil {
				handleBadFormSubmission(c)
				return
			}

			if err := warehouse.AddLegacyProducts(s.State.DB, requestData.Products); err != nil {
				log.Println(err.Error())
				handleBadFormSubmission(c)
				return
			}

		} else {
			handleBadFormSubmission(c)
			return
		}

		c.Status(http.StatusOK)
		_ = views.SuccessUploadingView().Render(c.Writer)
	}
}

func handleBadFormSubmission(c *gin.Context) {
	c.Status(http.StatusBadRequest)
	_ = views.ErrorUploadingView().Render(c.Writer)
}
