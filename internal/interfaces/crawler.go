package interfaces

import "github.com/gin-gonic/gin"

type CrawlerControllerInterface interface {
	AddCrawler(ctx *gin.Context)
}

type CrawlerServiceInterface interface{
	AddCrawler(c *gin.Context)
}