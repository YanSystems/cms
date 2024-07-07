package repositories

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/YanSystems/cms/pkg/models"
	utils "github.com/YanSystems/cms/pkg/utils/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateContent(t *testing.T) {
	testsCollection := uuid.New().String()

	client, err := utils.ConnectToDB("./../../../.env")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Database("content").Collection(testsCollection).Drop(context.TODO()); err != nil {
			log.Fatal(err)
		}
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repo := ContentRepository{
		DB: client.Database("content"),
	}

	initialContent := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "Initial Title",
		Description: "Initial Description",
		Body:        "Initial Body",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	// Insert initial content for testing update
	_, err = repo.CreateContent(testsCollection, initialContent)
	assert.NoError(t, err)

	defer func() {
		_, err := repo.DeleteContent(testsCollection, initialContent.Id)
		assert.NoError(t, err)
	}()

	t.Run("Successful Update", func(t *testing.T) {
		updatedContent := &models.Content{
			Class:       "updated-class",
			Title:       "Updated Title",
			Description: "Updated Description",
			Body:        "Updated Body",
			IsPublic:    false,
			Views:       100,
			CreatorId:   uuid.New().String(),
			UpdatedAt:   time.Now(),
			CreatedAt:   initialContent.CreatedAt,
		}

		id, err := repo.UpdateContent(testsCollection, initialContent.Id, updatedContent)
		assert.NoError(t, err)
		assert.Equal(t, initialContent.Id, id)

		content, err := repo.GetContent(testsCollection, id)
		assert.NoError(t, err)
		assert.Equal(t, updatedContent.Class, content.Class)
		assert.Equal(t, updatedContent.Title, content.Title)
		assert.Equal(t, updatedContent.Description, content.Description)
		assert.Equal(t, updatedContent.Body, content.Body)
		assert.Equal(t, updatedContent.IsPublic, content.IsPublic)
		assert.Equal(t, updatedContent.Views, content.Views)
		assert.Equal(t, updatedContent.CreatorId, content.CreatorId)
		assert.WithinDuration(t, updatedContent.UpdatedAt.UTC(), content.UpdatedAt.UTC(), time.Second)
		assert.WithinDuration(t, initialContent.CreatedAt.UTC(), content.CreatedAt.UTC(), time.Second)
	})

	// t.Run("Partial Update", func(t *testing.T) {
	// 	partialUpdatedContent := &models.Content{
	// 		Title: "Partially Updated Title",
	// 	}

	// 	id, err := repo.UpdateContent(testsCollection, initialContent.Id, partialUpdatedContent)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, initialContent.Id, id)

	// 	// Retrieve the updated content
	// 	content, err := repo.GetContent(testsCollection, id)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, initialContent.Class, content.Class)
	// 	assert.Equal(t, partialUpdatedContent.Title, content.Title)
	// 	assert.Equal(t, initialContent.Description, content.Description)
	// 	assert.Equal(t, initialContent.Body, content.Body)
	// 	assert.Equal(t, initialContent.IsPublic, content.IsPublic)
	// 	assert.Equal(t, initialContent.Views, content.Views)
	// 	assert.Equal(t, initialContent.CreatorId, content.CreatorId)
	// 	assert.WithinDuration(t, initialContent.UpdatedAt, content.UpdatedAt, time.Second)
	// 	assert.Equal(t, initialContent.CreatedAt, content.CreatedAt)
	// })

	// t.Run("Non-Existent Content", func(t *testing.T) {
	// 	nonExistentId := uuid.New().String()
	// 	updatedContent := &models.Content{
	// 		Class: "non-existent-class",
	// 	}

	// 	id, err := repo.UpdateContent(testsCollection, nonExistentId, updatedContent)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, "", id)
	// })

	// t.Run("Database Connection Nil", func(t *testing.T) {
	// 	repoNilDB := ContentRepository{DB: nil}
	// 	updatedContent := &models.Content{
	// 		Class: "updated-class",
	// 	}

	// 	id, err := repoNilDB.UpdateContent(testsCollection, initialContent.Id, updatedContent)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, "database connection is nil", err.Error())
	// 	assert.Equal(t, "", id)
	// })

	// t.Run("Invalid Content ID", func(t *testing.T) {
	// 	invalidId := "invalid-uuid"
	// 	updatedContent := &models.Content{
	// 		Class: "updated-class",
	// 	}

	// 	id, err := repo.UpdateContent(testsCollection, invalidId, updatedContent)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, "", id)
	// })
}
