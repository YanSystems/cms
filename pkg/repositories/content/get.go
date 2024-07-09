package repositories

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/YanSystems/cms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *ContentRepository) GetContent(coll string, id string) (*models.ReadContent, error) {
	slog.Debug("GetContent called", "collection", coll, "id", id)
	if r.DB == nil {
		err := errors.New("database connection is nil")
		slog.Error("Database connection is nil", "error", err)
		return nil, err
	}

	var content models.ReadContent

	err := r.DB.Collection(coll).FindOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
	).Decode(&content)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			err := errors.New("content not found")
			slog.Error("Content not found", "collection", coll, "id", id, "error", err)
			return nil, err
		}
		slog.Error("Failed to find content", "collection", coll, "id", id, "error", err)
		return nil, err
	}

	slog.Info("Content retrieved successfully", "content", content)
	return &content, nil
}

func (r *ContentRepository) GetCollection(coll string) ([]models.Content, error) {
	slog.Debug("GetCollection called", "collection", coll)

	results, err := r.DB.Collection(coll).Find(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		slog.Error("Failed to find collection", "collection", coll, "error", err)
		return nil, err
	}

	var contents []models.Content
	err = results.All(context.TODO(), &contents)
	if err != nil {
		err := fmt.Errorf("failed to decode results: %s", err.Error())
		slog.Error("Failed to decode results", "collection", coll, "error", err)
		return nil, err
	}

	if len(contents) == 0 {
		slog.Info("No contents found in collection", "collection", coll)
		return []models.Content{}, nil
	}

	slog.Info("Collection retrieved successfully", "collection", coll, "contents", contents)
	return contents, nil
}

func (r *ContentRepository) GetClass(coll string, class string) ([]models.Content, error) {
	slog.Debug("GetClass called", "collection", coll, "class", class)

	results, err := r.DB.Collection(coll).Find(
		context.TODO(),
		bson.D{{Key: "class", Value: class}},
	)
	if err != nil {
		slog.Error("Failed to find class contents", "collection", coll, "class", class, "error", err)
		return nil, err
	}

	var contents []models.Content
	err = results.All(context.TODO(), &contents)
	if err != nil {
		err := fmt.Errorf("failed to decode results: %s", err.Error())
		slog.Error("Failed to decode class contents", "collection", coll, "class", class, "error", err)
		return nil, err
	}

	if len(contents) == 0 {
		slog.Info("No contents found for class", "collection", coll, "class", class)
		return []models.Content{}, nil
	}

	slog.Info("Class contents retrieved successfully", "collection", coll, "class", class, "contents", contents)
	return contents, nil
}
