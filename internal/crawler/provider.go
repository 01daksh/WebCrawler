package crawler

import (
	"WebCrawler/internal/interfaces"

	"github.com/google/wire"
)


var ControllerSet = wire.NewSet(
	NewCrawlerController,
	NewCrawlerService,

	wire.Bind(new(interfaces.CrawlerControllerInterface), new(*CrawlerController)),
	wire.Bind(new(interfaces.CrawlerServiceInterface), new(*CrawlerService)),
)
