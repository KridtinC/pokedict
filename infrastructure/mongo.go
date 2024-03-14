package infrastructure

import (
	"context"

	"github.com/KridtinC/pokedict/config"
	"github.com/KridtinC/pokedict/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Get().DatabaseConfig.ConnectionString))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	logger.Get().Info("Connected to MongoDB")
	return client.Database(config.Get().DatabaseConfig.Database)
}
