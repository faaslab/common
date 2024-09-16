package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/faaslab/common/logger"
)

func NewClient(ctx context.Context) (*mongo.Client, error) {
	// Implement the logic to create a function
	clientOptions := options.Client().ApplyURI("mongodb://faas:Admin123@localhost:27017/?retryWrites=" +
		"true&loadBalanced=false&serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=faas&authMechanism=SCRAM-SHA-256&directConnection=true")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("Error connecting to MongoDB: %v", err)
		return nil, err // Handle connection error
	}
	defer client.Disconnect(ctx)

	return client, nil
}
