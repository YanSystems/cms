package repositories

import (
	"context"
	"time"

	"github.com/YanSystems/cms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) UpdateContent(coll string, id string, updatedContent *models.UpdateContent) (string, error) {
	// Fetch the current content from the database
	currentContent, err := r.GetContent(coll, id)
	if err != nil {
		return "", err
	}

	// Update fields only if they are set in updatedContent
	if updatedContent.Class != nil {
		currentContent.Class = *updatedContent.Class
	}
	if updatedContent.Title != nil {
		currentContent.Title = *updatedContent.Title
	}
	if updatedContent.Description != nil {
		currentContent.Description = *updatedContent.Description
	}
	if updatedContent.Body != nil {
		currentContent.Body = *updatedContent.Body
	}
	if updatedContent.Views != nil {
		currentContent.Views = *updatedContent.Views
	}
	if updatedContent.CreatorId != nil {
		currentContent.CreatorId = *updatedContent.CreatorId
	}
	if updatedContent.IsPublic != nil {
		currentContent.IsPublic = *updatedContent.IsPublic
	}

	// Always update UpdatedAt to the current time
	currentContent.UpdatedAt = time.Now().UTC()

	// Ensure CreatedAt is never updated
	updatedContent.CreatedAt = &currentContent.CreatedAt

	// Update the document in the database
	_, err = r.DB.Collection(coll).UpdateOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
		bson.D{{Key: "$set", Value: currentContent}},
	)

	if err != nil {
		return "", err
	}

	return id, nil
}
