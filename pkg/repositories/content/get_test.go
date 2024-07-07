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

func TestGetContent(t *testing.T) {
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

	defer func() {
		_, err := repo.DeleteCollection(testsCollection)
		assert.NoError(t, err)
	}()

	initialContent := &models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "",
		Description: "",
		Body:        "",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	_, err = repo.CreateContent(testsCollection, initialContent)
	assert.NoError(t, err)

	defer func() {
		_, err := repo.DeleteContent(testsCollection, initialContent.Id)
		assert.NoError(t, err)
	}()

	t.Run("Successful Retrieval", func(t *testing.T) {
		content, err := repo.GetContent(testsCollection, initialContent.Id)
		assert.NoError(t, err)
		assert.NotNil(t, content)
		assert.Equal(t, initialContent.Id, content.Id)
		assert.Equal(t, initialContent.Class, content.Class)
		assert.Equal(t, initialContent.Title, content.Title)
		assert.Equal(t, initialContent.Description, content.Description)
		assert.Equal(t, initialContent.Body, content.Body)
		assert.Equal(t, initialContent.IsPublic, content.IsPublic)
		assert.Equal(t, initialContent.Views, content.Views)
		assert.Equal(t, initialContent.CreatorId, content.CreatorId)
		assert.WithinDuration(t, initialContent.UpdatedAt, content.UpdatedAt, time.Second)
		assert.WithinDuration(t, initialContent.CreatedAt, content.CreatedAt, time.Second)
	})

	t.Run("Non-Existent Content", func(t *testing.T) {
		nonExistentId := uuid.New().String()
		content, err := repo.GetContent(testsCollection, nonExistentId)
		assert.Error(t, err)
		assert.Equal(t, "content not found", err.Error())
		assert.Nil(t, content)
	})

	t.Run("Invalid Content ID", func(t *testing.T) {
		invalidId := "invalid-uuid"
		content, err := repo.GetContent(testsCollection, invalidId)
		assert.Error(t, err)
		assert.Nil(t, content)
	})

	t.Run("Database Connection Nil", func(t *testing.T) {
		repoNilDB := ContentRepository{DB: nil}
		content, err := repoNilDB.GetContent(testsCollection, initialContent.Id)
		assert.Error(t, err)
		assert.Equal(t, "database connection is nil", err.Error())
		assert.Nil(t, content)
	})
}

func TestGetCollection(t *testing.T) {
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
		Title:       "", // Title can be empty
		Description: "", // Description can be empty
		Body:        "", // Body can be empty
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	// Insert initial content for testing retrieval
	_, err = repo.CreateContent(testsCollection, initialContent)
	assert.NoError(t, err)

	defer func() {
		_, err := repo.DeleteContent(testsCollection, initialContent.Id)
		assert.NoError(t, err)
	}()

	t.Run("Successful Retrieval", func(t *testing.T) {
		contents, err := repo.GetCollection(testsCollection)
		assert.NoError(t, err)
		assert.NotNil(t, contents)
		assert.Len(t, contents, 1)

		content := contents[0]
		assert.Equal(t, initialContent.Id, content.Id)
		assert.Equal(t, initialContent.Class, content.Class)
		assert.Equal(t, initialContent.Title, content.Title)
		assert.Equal(t, initialContent.Description, content.Description)
		assert.Equal(t, initialContent.Body, content.Body)
		assert.Equal(t, initialContent.IsPublic, content.IsPublic)
		assert.Equal(t, initialContent.Views, content.Views)
		assert.Equal(t, initialContent.CreatorId, content.CreatorId)
		assert.WithinDuration(t, initialContent.UpdatedAt, content.UpdatedAt, time.Second)
		assert.WithinDuration(t, initialContent.CreatedAt, content.CreatedAt, time.Second)
	})

	t.Run("Empty Collection", func(t *testing.T) {
		emptyCollection := uuid.New().String()

		contents, err := repo.GetCollection(emptyCollection)
		assert.NoError(t, err)
		assert.NotNil(t, contents)
		assert.Len(t, contents, 0)
	})

	t.Run("Invalid collection", func(t *testing.T) {
		invalidCollection := ""
		contents, err := repo.GetCollection(invalidCollection)
		assert.Error(t, err)
		assert.Nil(t, contents)
	})
}

func TestGetClass(t *testing.T) {
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
		Title:       "",
		Description: "",
		Body:        "",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	_, err = repo.CreateContent(testsCollection, initialContent)
	assert.NoError(t, err)

	defer func() {
		_, err := repo.DeleteContent(testsCollection, initialContent.Id)
		assert.NoError(t, err)
	}()

	t.Run("Successful Retrieval", func(t *testing.T) {
		contents, err := repo.GetClass(testsCollection, "test-class")
		assert.NoError(t, err)
		assert.NotNil(t, contents)
		assert.Len(t, contents, 1)

		content := contents[0]
		assert.Equal(t, initialContent.Id, content.Id)
		assert.Equal(t, initialContent.Class, content.Class)
		assert.Equal(t, initialContent.Title, content.Title)
		assert.Equal(t, initialContent.Description, content.Description)
		assert.Equal(t, initialContent.Body, content.Body)
		assert.Equal(t, initialContent.IsPublic, content.IsPublic)
		assert.Equal(t, initialContent.Views, content.Views)
		assert.Equal(t, initialContent.CreatorId, content.CreatorId)
		assert.WithinDuration(t, initialContent.UpdatedAt, content.UpdatedAt, time.Second)
		assert.WithinDuration(t, initialContent.CreatedAt, content.CreatedAt, time.Second)
	})

	t.Run("Empty Class", func(t *testing.T) {
		emptyClass := "non-existent-class"

		contents, err := repo.GetClass(testsCollection, emptyClass)
		assert.NoError(t, err)
		assert.NotNil(t, contents)
		assert.Len(t, contents, 0)
	})
	t.Run("Database Operation Failure", func(t *testing.T) {
		invalidCollection := ""
		contents, err := repo.GetClass(invalidCollection, "test-class")
		assert.Error(t, err)
		assert.Nil(t, contents)
	})
}
