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
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateContent(t *testing.T) {
	testsCollection := uuid.New().String()

	client, err := utils.ConnectToDB("./../../.env")
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
		if err := repo.DB.Collection(testsCollection).Drop(context.TODO()); err != nil {
			log.Fatal(err)
		}
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

		var insertedContent models.Content
		err = repo.DB.Collection(testsCollection).FindOne(context.TODO(), bson.M{"id": content.Id}).Decode(&insertedContent)
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

		_, err = repo.DB.Collection(testsCollection).DeleteOne(context.TODO(), bson.M{"id": content.Id})
		assert.NoError(t, err)
	})

	t.Run("Duplicate Ids", func(t *testing.T) {
		_, err = repo.CreateContent(testsCollection, content)
		assert.NoError(t, err)
		_, err = repo.CreateContent(testsCollection, content)
		assert.Error(t, err)

		_, err = repo.DB.Collection(testsCollection).DeleteOne(context.TODO(), bson.M{"id": content.Id})
		assert.NoError(t, err)
	})

	t.Run("Non-UUID Id", func(*testing.T) {
		contentWithNonUUIDID := *content
		contentWithNonUUIDID.Id = "non-uuid-string"

		id, err := repo.CreateContent(testsCollection, &contentWithNonUUIDID)
		assert.NoError(t, err)
		assert.Equal(t, contentWithNonUUIDID.Id, id)

		var insertedContent models.Content
		err = repo.DB.Collection(testsCollection).FindOne(context.TODO(), bson.M{"id": contentWithNonUUIDID.Id}).Decode(&insertedContent)
		assert.NoError(t, err)
		assert.Equal(t, contentWithNonUUIDID.Id, insertedContent.Id)
		assert.Equal(t, contentWithNonUUIDID.Class, insertedContent.Class)
		assert.Equal(t, contentWithNonUUIDID.Title, insertedContent.Title)
		assert.Equal(t, contentWithNonUUIDID.Description, insertedContent.Description)
		assert.Equal(t, contentWithNonUUIDID.Body, insertedContent.Body)
		assert.Equal(t, contentWithNonUUIDID.IsPublic, insertedContent.IsPublic)
		assert.Equal(t, contentWithNonUUIDID.Views, insertedContent.Views)
		assert.Equal(t, contentWithNonUUIDID.CreatorId, insertedContent.CreatorId)
		assert.WithinDuration(t, contentWithNonUUIDID.UpdatedAt, insertedContent.UpdatedAt, time.Second)
		assert.WithinDuration(t, contentWithNonUUIDID.CreatedAt, insertedContent.CreatedAt, time.Second)

		_, err = repo.DB.Collection(testsCollection).DeleteOne(context.TODO(), bson.M{"id": contentWithNonUUIDID.Id})
		assert.NoError(t, err)
	})
}
