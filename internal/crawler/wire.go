//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package crawler

import "github.com/google/wire"

func NewCrawlerWire() *CrawlerController{
	wire.Build(
		ControllerSet,
	)

	return &CrawlerController{}
}