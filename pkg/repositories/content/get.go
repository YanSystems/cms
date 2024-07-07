package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/YanSystems/cms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *ContentRepository) GetContent(coll string, id string) (*models.Content, error) {
	if r.DB == nil {
		return nil, errors.New("database connection is nil")
	}

	var content models.Content

	err := r.DB.Collection(coll).FindOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
	).Decode(&content)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("content not found")
		}
		return nil, err
	}

	return &content, nil
}

func (r *ContentRepository) GetCollection(coll string) ([]models.Content, error) {
	results, err := r.DB.Collection(coll).Find(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		return nil, err
	}

	var contents []models.Content
	err = results.All(context.TODO(), &contents)
	if err != nil {
		return nil, fmt.Errorf("failed to decode results: %s", err.Error())
	}

	if len(contents) == 0 {
		return []models.Content{}, nil
	}

	return contents, nil
}

func (r *ContentRepository) GetClass(coll string, class string) ([]models.Content, error) {
	results, err := r.DB.Collection(coll).Find(
		context.TODO(),
		bson.D{{Key: "class", Value: class}},
	)
	if err != nil {
		return nil, err
	}

	var contents []models.Content
	err = results.All(context.TODO(), &contents)
	if err != nil {
		return nil, fmt.Errorf("failed to decode results: %s", err.Error())
	}

	if len(contents) == 0 {
		return []models.Content{}, nil
	}

	return contents, nil
}
