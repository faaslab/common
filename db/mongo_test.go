package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	var localURI = "mongodb://faas:Admin123@localhost:27017/?retryWrites=" +
		"true&loadBalanced=false&serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=faas&authMechanism=SCRAM-SHA-256&directConnection=true"

	_, err := NewClient(context.Background(), localURI)
	assert.NoError(t, err)
}
