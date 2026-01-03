package main

import (
	"WebCrawler/configs"

	"WebCrawler/appinit"
	"WebCrawler/internal/crawler"

	core "github.com/01daksh/crawler-core"
	"github.com/gin-gonic/gin"

	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configs.InitializeConfig()
	ctx, stop := signal.NotifyContext(
		context.TODO(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	runHttpServer(ctx)
}

func runHttpServer(ctx context.Context) {

	router := gin.Default()

	appinit.Initialize(ctx)
	router.Use(core.RequestIdMiddleWare)
	router.POST("/crawl", crawler.NewCrawlerWire().AddCrawler)
	// fmt.Printf("%s\t", viper.GetString("db.mongoDB.mongoURI"))
	router.Run()
}
