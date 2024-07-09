package repositories

import (
	"context"
	"log/slog"
	"time"

	"github.com/YanSystems/cms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) UpdateContent(coll string, id string, updatedContent *models.UpdateContent) (string, error) {
	slog.Info("UpdateContent called", "collection", coll, "id", id)

	// Fetch the current content from the database
	currentContent, err := r.GetContent(coll, id)
	if err != nil {
		slog.Error("Failed to get current content", "collection", coll, "id", id, "error", err)
		return "", err
	}
	slog.Debug("Current content fetched", "currentContent", currentContent)

	// Update fields only if they are set in updatedContent
	if updatedContent.Class != nil {
		slog.Debug("Updating Class", "oldValue", currentContent.Class, "newValue", *updatedContent.Class)
		currentContent.Class = *updatedContent.Class
	}
	if updatedContent.Title != nil {
		slog.Debug("Updating Title", "oldValue", currentContent.Title, "newValue", *updatedContent.Title)
		currentContent.Title = *updatedContent.Title
	}
	if updatedContent.Description != nil {
		slog.Debug("Updating Description", "oldValue", currentContent.Description, "newValue", *updatedContent.Description)
		currentContent.Description = *updatedContent.Description
	}
	if updatedContent.Body != nil {
		slog.Debug("Updating Body", "oldValue", currentContent.Body, "newValue", *updatedContent.Body)
		currentContent.Body = *updatedContent.Body
	}
	if updatedContent.Views != nil {
		slog.Debug("Updating Views", "oldValue", currentContent.Views, "newValue", *updatedContent.Views)
		currentContent.Views = *updatedContent.Views
	}
	if updatedContent.CreatorId != nil {
		slog.Debug("Updating CreatorId", "oldValue", currentContent.CreatorId, "newValue", *updatedContent.CreatorId)
		currentContent.CreatorId = *updatedContent.CreatorId
	}
	if updatedContent.IsPublic != nil {
		slog.Debug("Updating IsPublic", "oldValue", currentContent.IsPublic, "newValue", *updatedContent.IsPublic)
		currentContent.IsPublic = *updatedContent.IsPublic
	}

	// Always update UpdatedAt to the current time
	currentContent.UpdatedAt = time.Now().UTC()
	slog.Debug("Updated UpdatedAt", "newValue", currentContent.UpdatedAt)

	// Ensure CreatedAt is never updated
	updatedContent.CreatedAt = &currentContent.CreatedAt

	// Update the document in the database
	_, err = r.DB.Collection(coll).UpdateOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
		bson.D{{Key: "$set", Value: currentContent}},
	)

	if err != nil {
		slog.Error("Failed to update content", "collection", coll, "id", id, "error", err)
		return "", err
	}

	slog.Info("Content updated successfully", "collection", coll, "id", id)
	return id, nil
}
