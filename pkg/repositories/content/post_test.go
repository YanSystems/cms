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

func TestCreateContent(t *testing.T) {
	testsCollection := uuid.New().String()

	client, err := utils.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repo := ContentRepository{
		DB: client.Database("content"),
	}

	defer func() {
		_, err := repo.DeleteCollection(testsCollection)
		assert.NoError(t, err)
	}()

	content := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "test-title",
		Description: "test-description",
		Body:        "test-body",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	t.Run("Successful Creation", func(t *testing.T) {
		id, err := repo.CreateContent(testsCollection, content)
		assert.NoError(t, err)
		assert.Equal(t, content.Id, id)

		insertedContent, err := repo.GetContent(testsCollection, id)
		assert.NoError(t, err)
		assert.Equal(t, content.Id, insertedContent.Id)
		assert.Equal(t, content.Class, insertedContent.Class)
		assert.Equal(t, content.Title, insertedContent.Title)
		assert.Equal(t, content.Description, insertedContent.Description)
		assert.Equal(t, content.Body, insertedContent.Body)
		assert.Equal(t, content.IsPublic, insertedContent.IsPublic)
		assert.Equal(t, content.Views, insertedContent.Views)
		assert.Equal(t, content.CreatorId, insertedContent.CreatorId)
		assert.WithinDuration(t, content.UpdatedAt, insertedContent.UpdatedAt, time.Second)
		assert.WithinDuration(t, content.CreatedAt, insertedContent.CreatedAt, time.Second)

		deletedId, err := repo.DeleteContent(testsCollection, id)
		assert.NoError(t, err)
		assert.Equal(t, id, deletedId)
	})

	t.Run("Unpopulated or incorrect fields", func(t *testing.T) {
		testCases := []struct {
			name         string
			unpopulate   func(content *models.Content)
			missingField string
		}{
			{
				name: "Missing or non-UUID Id",
				unpopulate: func(content *models.Content) {
					content.Id = ""
				},
				missingField: "Id",
			},
			{
				name: "Missing or empty Class",
				unpopulate: func(content *models.Content) {
					content.Class = ""
				},
				missingField: "Class",
			},
			{
				name: "Missing or non-UUID CreatorId",
				unpopulate: func(content *models.Content) {
					content.CreatorId = ""
				},
				missingField: "CreatorId",
			},
			{
				name: "Missing UpdatedAt",
				unpopulate: func(content *models.Content) {
					content.UpdatedAt = time.Time{}
				},
				missingField: "UpdatedAt",
			},
			{
				name: "Missing CreatedAt",
				unpopulate: func(content *models.Content) {
					content.CreatedAt = time.Time{}
				},
				missingField: "CreatedAt",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				contentClone := *content
				tc.unpopulate(&contentClone)

				_, err := repo.CreateContent(testsCollection, &contentClone)
				assert.Error(t, err, "Expected error for missing field: %s", tc.missingField)
			})
		}
	})

	t.Run("Duplicate Ids", func(t *testing.T) {
		id, err := repo.CreateContent(testsCollection, content)
		assert.NoError(t, err)
		_, err = repo.CreateContent(testsCollection, content)
		assert.Error(t, err)

		deletedId, err := repo.DeleteContent(testsCollection, id)
		assert.NoError(t, err)
		assert.Equal(t, id, deletedId)
	})
}
