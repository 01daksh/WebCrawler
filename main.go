package main

import (
	"WebCrawler/configs"

	"WebCrawler/appinit"
	"WebCrawler/internal/crawler"

	"log"

	core "github.com/01daksh/crawler-core"
	"github.com/01daksh/crawler-core/database/provider"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v", r)
			// Attempt to disconnect MongoDB client on panic
			ctx := context.Background()
			client, err := provider.GetTenantMongoDb().GetClient()
			if err != nil {
				log.Printf("Error getting MongoDB client on panic: %v", err)
				os.Exit(1)
			}
			if err := client.Disconnect(ctx); err != nil {
				log.Printf("Error disconnecting MongoDB on panic: %v", err)
			}
			os.Exit(1)
		}
	}()

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

	// Create HTTP server with proper configuration
	srv := &http.Server{
		Addr:    viper.GetString("server.httpServerAddress"),
		Handler: router,
	}

	// Channel to signal server shutdown completion
	done := make(chan bool, 1)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting HTTP server on :8086")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
		done <- true
	}()

	// Wait for context cancellation (signal)
	<-ctx.Done()
	log.Printf("Shutting down server...")

	// Create a deadline for shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer shutdownCancel()

	// Attempt graceful shutdown of HTTP server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server forced to shutdown: %v", err)
	}

	// Wait for server to finish
	<-done
	log.Printf("HTTP server stopped")

	// Disconnect MongoDB client
	client, err := provider.GetTenantMongoDb().GetClient()
	if err != nil {
		log.Printf("Error getting MongoDB client on shutdown: %v", err)
		os.Exit(1)
	}
	if err := client.Disconnect(shutdownCtx); err != nil {
		log.Printf("Error disconnecting MongoDB: %v", err)
	}

	log.Printf("Graceful shutdown completed")
}
