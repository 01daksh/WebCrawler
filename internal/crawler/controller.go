package crawler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type CrawlerController struct{

}


func (c *CrawlerController) AddCrawler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "reached",
	})

	
}