package repositories

import (
	"context"

	"github.com/YanSystems/cms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) UpdateContent(coll string, id string, updatedContent *models.Content) (string, error) {
	update := bson.M{}
	if updatedContent.Class != "" {
		update["class"] = updatedContent.Class
	}
	if updatedContent.Title != "" {
		update["title"] = updatedContent.Title
	}
	if updatedContent.Description != "" {
		update["description"] = updatedContent.Description
	}
	if updatedContent.Body != "" {
		update["body"] = updatedContent.Body
	}
	// Ensure IsPublic is always updated
	update["is_public"] = updatedContent.IsPublic
	if updatedContent.Views != 0 {
		update["views"] = updatedContent.Views
	}
	if updatedContent.CreatorId != "" {
		update["creator_id"] = updatedContent.CreatorId
	}
	if !updatedContent.UpdatedAt.IsZero() {
		update["updated_at"] = updatedContent.UpdatedAt
	}
	if !updatedContent.CreatedAt.IsZero() {
		update["created_at"] = updatedContent.CreatedAt
	}

	_, err := r.DB.Collection(coll).UpdateOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
		bson.D{{Key: "$set", Value: update}},
	)

	if err != nil {
		return "", err
	}

	return id, nil
}
