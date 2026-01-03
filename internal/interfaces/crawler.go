package interfaces

import (
	"WebCrawler/internal/models"
	"context"
	"WebCrawler/common"

	"github.com/gin-gonic/gin"
)

type CrawlerControllerInterface interface {
	AddCrawler(ctx *gin.Context)
}

type CrawlerServiceInterface interface {
	AddCrawler(ctx context.Context, req models.CrawlerBO) ([]common.LinkInformation, error)
}

type CrawlerRepoInterface interface {
	AddCrawler(ctx context.Context,links []common.LinkInformation) error
}
