package crawler

import (
	"WebCrawler/internal/interfaces"
	"WebCrawler/internal/models"
	"net/http"

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

func (c *CrawlerController) AddCrawler(gctx *gin.Context) {
	var req models.CrawlerBO
	err := gctx.ShouldBindJSON(&req)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
	}
	ctx := gctx.Request.Context()

	links, err := c.crawServ.AddCrawler(ctx, req)
	if err != nil{
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	gctx.JSON(http.StatusOK, links)
}
