package db

import (
	"context"
	"time"

	"github.com/japnoor/daelog/pkg/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect(uri string) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal("failed to connect to MongoDB", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Fatal("failed to ping MongoDB", err)
	}

	logger.Info("connected to MongoDB").Send()

	return &MongoDB{
		Client:   client,
		Database: client.Database("daelog"),
	}
}
