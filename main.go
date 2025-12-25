package main

import (
	core "github.com/01daksh/crawler-core"
	"github.com/gin-gonic/gin"
)

func main() {
	runHttpServer()
}

func runHttpServer() {
	router := gin.Default()

	router.Use(core.RequestIdMiddleWare)
	router.POST("/crawl", )

	router.Run()
}
