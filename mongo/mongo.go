// mongo.go
package mongo

import (
	"context"
	"time"

	"github.com/pratyakshroy47/gql-go/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	*mongo.Client
}

func NewClient(uri string) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB")
	return &Client{client}, nil
}

func (c *Client) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.Disconnect(ctx)
	if err != nil {
		logger.Error("Failed to disconnect from MongoDB", err)
	} else {
		logger.Info("Disconnected from MongoDB")
	}
}