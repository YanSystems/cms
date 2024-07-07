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

func TestDeleteContent(t *testing.T) {
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

	_, err = repo.CreateContent(testsCollection, initialContent)
	assert.NoError(t, err)

	t.Run("Successful Deletion", func(t *testing.T) {
		id, err := repo.DeleteContent(testsCollection, initialContent.Id)
		assert.NoError(t, err)
		assert.Equal(t, initialContent.Id, id)

		content, err := repo.GetContent(testsCollection, id)
		assert.Error(t, err)
		assert.Equal(t, "content not found", err.Error())
		assert.Nil(t, content)
	})

	t.Run("Non-Existent Content", func(t *testing.T) {
		nonExistentId := uuid.New().String()
		id, err := repo.DeleteContent(testsCollection, nonExistentId)
		assert.NoError(t, err)
		assert.Equal(t, nonExistentId, id)
	})

	t.Run("Invalid Content ID", func(t *testing.T) {
		invalidId := "invalid-uuid"
		id, err := repo.DeleteContent(testsCollection, invalidId)
		assert.NoError(t, err)
		assert.Equal(t, invalidId, id)
	})
}

func TestDeleteClass(t *testing.T) {
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

	initialContent1 := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "Initial Title 1",
		Description: "Initial Description 1",
		Body:        "Initial Body 1",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	initialContent2 := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "Initial Title 2",
		Description: "Initial Description 2",
		Body:        "Initial Body 2",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	_, err = repo.CreateContent(testsCollection, initialContent1)
	assert.NoError(t, err)
	_, err = repo.CreateContent(testsCollection, initialContent2)
	assert.NoError(t, err)

	t.Run("Successful Deletion", func(t *testing.T) {
		ids, err := repo.DeleteClass(testsCollection, "test-class")
		assert.NoError(t, err)
		assert.Len(t, ids, 2)
		assert.Contains(t, ids, initialContent1.Id)
		assert.Contains(t, ids, initialContent2.Id)

		for _, id := range ids {
			content, err := repo.GetContent(testsCollection, id)
			assert.Error(t, err)
			assert.Equal(t, "content not found", err.Error())
			assert.Nil(t, content)
		}
	})

	t.Run("Non-Existent Class", func(t *testing.T) {
		ids, err := repo.DeleteClass(testsCollection, "non-existent-class")
		assert.NoError(t, err)
		assert.Len(t, ids, 0)
	})
	t.Run("Invalid Collection", func(t *testing.T) {
		invalidCollection := ""
		ids, err := repo.DeleteClass(invalidCollection, "test-class")
		assert.Error(t, err)
		assert.Nil(t, ids)
	})
}

func TestDeleteCollection(t *testing.T) {
	testsCollection := uuid.New().String()

	client, err := utils.ConnectToDB("./../../../.env")
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

	initialContent1 := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "Initial Title 1",
		Description: "Initial Description 1",
		Body:        "Initial Body 1",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	initialContent2 := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "Initial Title 2",
		Description: "Initial Description 2",
		Body:        "Initial Body 2",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	_, err = repo.CreateContent(testsCollection, initialContent1)
	assert.NoError(t, err)
	_, err = repo.CreateContent(testsCollection, initialContent2)
	assert.NoError(t, err)

	t.Run("Successful Deletion", func(t *testing.T) {
		ids, err := repo.DeleteCollection(testsCollection)
		assert.NoError(t, err)
		assert.Len(t, ids, 2)
		assert.Contains(t, ids, initialContent1.Id)
		assert.Contains(t, ids, initialContent2.Id)

		collections, err := client.Database("content").ListCollectionNames(context.TODO(), bson.D{})
		assert.NoError(t, err)
		assert.NotContains(t, collections, testsCollection)
	})

	t.Run("Empty Collection", func(t *testing.T) {
		emptyCollection := uuid.New().String()
		repo.CreateContent(emptyCollection, initialContent1)
		repo.DeleteContent(emptyCollection, initialContent1.Id)
		assert.NoError(t, err)

		ids, err := repo.DeleteCollection(emptyCollection)
		assert.NoError(t, err)
		assert.Len(t, ids, 0)

		collections, err := client.Database("content").ListCollectionNames(context.TODO(), bson.D{})
		assert.NoError(t, err)
		assert.NotContains(t, collections, emptyCollection)
	})

	t.Run("Invalid Collection", func(t *testing.T) {
		invalidCollection := ""
		ids, err := repo.DeleteCollection(invalidCollection)
		assert.Error(t, err)
		assert.Nil(t, ids)
	})
}
