package crawler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrawlerService struct {
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{}
}

func (*CrawlerService) AddCrawler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "reached",
	})

}
