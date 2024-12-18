package mongodb

import (
	"context"
	"os"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_url := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_url))

	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB!")

	return client.Database(mongodb_database), nil
}
