package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Client, error) {
	slog.Debug("Loading environment variable 'YAN_CMS_DB_URI'...")
	uri := os.Getenv("YAN_CMS_DB_URI")
	if uri == "" {
		err := fmt.Errorf("set your 'YAN_CMS_DB_URI' environment variable")
		slog.Error("Environment variable YAN_CMS_DB_URI not set", "error", err)
		return nil, err
	}
	slog.Debug("Environment variable YAN_CMS_DB_URI loaded", "uri", uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "error", err)
		return nil, err
	}
	slog.Info("Successfully connected to MongoDB!")

	return client, nil
}
