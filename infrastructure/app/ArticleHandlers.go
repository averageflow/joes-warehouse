package app

import "github.com/gin-gonic/gin"

func (s *ApplicationServer) getArticlesHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addArticlesHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addArticlesFromLegacyFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) addArticlesFromFileHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) modifyArticleHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}

func (s *ApplicationServer) deleteArticleHandler() func(*gin.Context) {
	return func(c *gin.Context) {}
}
