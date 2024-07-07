package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(pathToEnv string) (*mongo.Client, error) {
	err := godotenv.Load(pathToEnv)
	if err != nil {
		panic(err)
	}
	uri := os.Getenv("DB_URI")
	if uri == "" {
		return nil, fmt.Errorf("set your 'DB_URI' environment variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}
