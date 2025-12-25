package interfaces

import "github.com/gin-gonic/gin"

type CrawlerInterface interface {
	AddCrawler(ctx *gin.Context)
}
