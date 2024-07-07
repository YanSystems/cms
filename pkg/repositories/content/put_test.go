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

	client, err := utils.ConnectToDB()
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

	t.Run("Successful Update", func(t *testing.T) {
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

		_, err := repo.CreateContent(testsCollection, initialContent)
		assert.NoError(t, err)

		defer func() {
			_, err := repo.DeleteContent(testsCollection, initialContent.Id)
			assert.NoError(t, err)
		}()

		newTitle := "Updated Title"
		newClass := "updated-class"
		newDescription := "Updated Description"
		newBody := "Updated Body"
		newIsPublic := false
		newViews := 100
		newCreatorId := uuid.New().String()
		updatedContent := &models.UpdateContent{
			Class:       &newClass,
			Title:       &newTitle,
			Description: &newDescription,
			Body:        &newBody,
			IsPublic:    &newIsPublic,
			Views:       &newViews,
			CreatorId:   &newCreatorId,
			UpdatedAt:   &initialContent.UpdatedAt,
			CreatedAt:   &initialContent.CreatedAt,
		}

		id, err := repo.UpdateContent(testsCollection, initialContent.Id, updatedContent)
		assert.NoError(t, err)
		assert.Equal(t, initialContent.Id, id)

		content, err := repo.GetContent(testsCollection, id)
		assert.NoError(t, err)
		assert.Equal(t, *updatedContent.Class, content.Class)
		assert.Equal(t, *updatedContent.Title, content.Title)
		assert.Equal(t, *updatedContent.Description, content.Description)
		assert.Equal(t, *updatedContent.Body, content.Body)
		assert.Equal(t, *updatedContent.IsPublic, content.IsPublic)
		assert.Equal(t, *updatedContent.Views, content.Views)
		assert.Equal(t, *updatedContent.CreatorId, content.CreatorId)
		assert.NotEqual(t, initialContent.UpdatedAt.UTC(), content.UpdatedAt.UTC())
		assert.WithinDuration(t, initialContent.CreatedAt.UTC(), content.CreatedAt.UTC(), time.Second)
	})

	t.Run("Partial Update", func(t *testing.T) {
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

		_, err := repo.CreateContent(testsCollection, initialContent)
		assert.NoError(t, err)

		defer func() {
			_, err := repo.DeleteContent(testsCollection, initialContent.Id)
			assert.NoError(t, err)
		}()

		newTitle := "Partially Updated Title"
		partialUpdatedContent := &models.UpdateContent{
			Title: &newTitle,
		}

		id, err := repo.UpdateContent(testsCollection, initialContent.Id, partialUpdatedContent)
		assert.NoError(t, err)
		assert.Equal(t, initialContent.Id, id)

		// Retrieve the updated content
		content, err := repo.GetContent(testsCollection, id)
		assert.NoError(t, err)
		assert.Equal(t, initialContent.Class, content.Class)
		assert.Equal(t, *partialUpdatedContent.Title, content.Title)
		assert.Equal(t, initialContent.Description, content.Description)
		assert.Equal(t, initialContent.Body, content.Body)
		assert.Equal(t, initialContent.IsPublic, content.IsPublic)
		assert.Equal(t, initialContent.Views, content.Views)
		assert.Equal(t, initialContent.CreatorId, content.CreatorId)
		assert.NotEqual(t, initialContent.UpdatedAt.UTC(), content.UpdatedAt.UTC())
		assert.WithinDuration(t, initialContent.CreatedAt.UTC(), content.CreatedAt.UTC(), time.Second)
	})

	t.Run("Non-Existent Content", func(t *testing.T) {
		nonExistentId := uuid.New().String()
		newClass := "non-existent-class"
		updatedContent := &models.UpdateContent{
			Class: &newClass,
		}

		_, err := repo.UpdateContent(testsCollection, nonExistentId, updatedContent)
		assert.Error(t, err)
	})

	t.Run("Invalid Content ID", func(t *testing.T) {
		invalidId := "invalid-uuid"
		newClass := "updated-class"
		updatedContent := &models.UpdateContent{
			Class: &newClass,
		}

		_, err := repo.UpdateContent(testsCollection, invalidId, updatedContent)
		assert.Error(t, err)
	})
}
