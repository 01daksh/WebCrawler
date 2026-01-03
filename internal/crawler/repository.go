package crawler

import (
	"WebCrawler/common"
	"context"
	"sync"

	"github.com/01daksh/crawler-core/database/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CrawlerRepository struct {
	indexOnce sync.Once
}

func NewCrawlerRepository() *CrawlerRepository {
	return &CrawlerRepository{}
}

func (r *CrawlerRepository) AddCrawler(
	ctx context.Context,
	links []common.LinkInformation,
) error {

	if len(links) == 0 {
		return nil
	}

	client, err := provider.GetTenantMongoDb().GetClient()
	if err != nil {
		return err
	}

	collection := client.
		Database("crawler").
		Collection("links")


	if err := r.ensureIndexes(ctx, collection); err != nil {
		return err
	}

	models := make([]mongo.WriteModel, 0, len(links))

	for _, link := range links {
		model := mongo.NewUpdateOneModel().
			SetFilter(bson.M{"url": link.Link}).
			SetUpdate(bson.M{"$setOnInsert": link}).
			SetUpsert(true)

		models = append(models, model)
	}

	opts := options.BulkWrite().SetOrdered(false)

	_, err = collection.BulkWrite(ctx, models, opts)
	return err
}


func (r *CrawlerRepository) ensureIndexes(
	ctx context.Context,
	collection *mongo.Collection,
) error {

	var err error

	r.indexOnce.Do(func() {
		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "url", Value: 1},
			},
			Options: options.Index().
				SetUnique(true),
		}

		_, err = collection.Indexes().CreateOne(ctx, indexModel)
	})

	return err
}
