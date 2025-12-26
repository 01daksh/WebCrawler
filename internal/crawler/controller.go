package crawler

import (
	"WebCrawler/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type CrawlerController struct {
	crawServ interfaces.CrawlerServiceInterface
}

func NewCrawlerController(serv interfaces.CrawlerServiceInterface) *CrawlerController {
	return &CrawlerController{
		crawServ: serv,
	}
}

func (c *CrawlerController) AddCrawler(ctx *gin.Context) {
	c.crawServ.AddCrawler(ctx)
}