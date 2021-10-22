package app

import "github.com/gin-gonic/gin"

func (s *ApplicationServer) getProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addProductsHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addProductsFromLegacyFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addProductsFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) modifyProductHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) deleteProductHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}
