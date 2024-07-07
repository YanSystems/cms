package repositories

import (
	"context"
	"errors"

	"github.com/YanSystems/cms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContentRepository struct {
	DB *mongo.Database
}

// create a new content document into a given collection
// returns the id of the created document
// requires that the passed in content has all fields populated
func (r *ContentRepository) CreateContent(coll string, content *models.Content) (string, error) {
	var existingContent models.Content
	err := r.DB.Collection(coll).FindOne(context.TODO(), bson.M{"id": content.Id}).Decode(&existingContent)
	if err == nil {
		return "", errors.New("content with this ID already exists")
	} else if err != mongo.ErrNoDocuments {
		return "", err
	}

	_, err = r.DB.Collection(coll).InsertOne(context.TODO(), content)
	if err != nil {
		return "", err
	}

	return content.Id, nil
}

// get content from a given collection by id
func (r *ContentRepository) GetContent(coll string, id string) (*models.Content, error) {
	return nil, nil
}

// get all contents	from a given collection
func (r *ContentRepository) GetCollection(coll string) ([]models.Content, error) {
	return nil, nil
}

// get all contents from a given collection that matches the given class
func (r *ContentRepository) GetClass(coll string, class string) ([]models.Content, error) {
	return nil, nil
}

// update a content in a given collection by id
// not all fields need to be populated
// a field that is not populated will not be modified.
// returns the id of the updated content
func (r *ContentRepository) UpdateContent(coll string, id string, updatedContent *models.Content) (string, error) {
	return "", nil
}

// delete content in a given collection by id
// returns the id of the deleted content
func (r *ContentRepository) DeleteContent(coll string, id string) (string, error) {
	return "", nil
}

// delete all contents in a given collection that matches the given class
// returns an array of ids of deleted contents
func (r *ContentRepository) DeleteClass(coll string, class string) ([]string, error) {
	return nil, nil
}

// delete all contents in a given collection
// returns an array of ids of deleted contents
func (r *ContentRepository) DeleteCollection(coll string) ([]string, error) {
	return nil, nil
}
