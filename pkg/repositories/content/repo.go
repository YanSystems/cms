package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ContentRepository struct {
	DB *mongo.Database
}
