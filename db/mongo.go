package db

import (
	"context"

	"github.com/faaslab/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, uri string) (*mongo.Client, error) {
	// local mongo
	// clientOptions := options.Client().ApplyURI(localURI)

	// atlas mongo
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Implement the logic to create a function
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Errorf("Error connecting to MongoDB: %v", err)
		return nil, err // Handle connection error
	}
	defer client.Disconnect(ctx)

	return client, nil
}
