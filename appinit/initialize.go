package appinit

import (
	"WebCrawler/common"
	"context"

	"github.com/01daksh/crawler-core/database/options"
	"github.com/01daksh/crawler-core/database/provider"
)

func Initialize(ctx context.Context) {
	initDB(ctx)
}

func initDB(ctx context.Context) {
	provider.GetTenantMongoDb().Initialize(ctx, options.WithMongoURI(common.GetMongoConnectionString(common.MongoUri)))
}
