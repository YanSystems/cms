package apitests

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/YanSystems/cms/pkg/models"
	repositories "github.com/YanSystems/cms/pkg/repositories/content"
	"github.com/YanSystems/cms/pkg/server"
	utils "github.com/YanSystems/cms/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func (chtc *ContentHandlerTestCase) RunContentDeleteTest(t *testing.T) {
	client, err := utils.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	api := server.Server{
		Port: "8000",
		DB:   client.Database("content"),
	}
	repo := repositories.ContentRepository{DB: api.DB}
	r := api.NewRouter()

	switch chtc.Case {
	case "DeleteContent":
		_, err := repo.DeleteContent(testsCollection, chtc.RequestPayload.Id)
		assert.NoError(t, err)
	case "DeleteCollection", "DeleteClass":
		contents := chtc.ArrayRequestPayload
		for _, content := range contents {
			content.Id = uuid.New().String()
			content.UpdatedAt = time.Now().UTC()
			content.CreatedAt = time.Now().UTC()
			_, err := repo.CreateContent(testsCollection, &content)
			assert.NoError(t, err)
		}
	}

	req, err := http.NewRequest("DELETE", baseUrl+chtc.Path, nil)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var responsePayload models.JsonResponse
	_ = json.NewDecoder(rr.Body).Decode(&responsePayload)

	assert.Equal(t, chtc.ExpectedStatus, rr.Code, "HTTP status code mismatch")
	assert.Equal(t, chtc.ExpectedResponse.Error, responsePayload.Error, "Error field mismatch")
	assert.Equal(t, chtc.ExpectedResponse.Message, responsePayload.Message, "Message field mismatch")
}

func TestHandleDeleteContentFound(t *testing.T) {
	content := models.Content{
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
	chtc := ContentHandlerTestCase{
		Case:           "DeleteContent",
		Path:           "/" + testsCollection + "/id/" + content.Id,
		RequestPayload: content,
		ExpectedStatus: http.StatusOK,
		ExpectedResponse: models.JsonResponse{
			Error:   false,
			Message: "Successfully deleted content",
		},
	}

	chtc.RunContentDeleteTest(t)
}

func TestHandleDeleteCollectionFound(t *testing.T) {
	contentOne := models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class1",
		Title:       "test-title1",
		Description: "test-description1",
		Body:        "test-body1",
		IsPublic:    false,
		Views:       10,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	contentTwo := models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class2",
		Title:       "test-title2",
		Description: "test-description2",
		Body:        "test-body2",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	var contents []models.Content
	contents = append(contents, contentOne, contentTwo)

	chtc := ContentHandlerTestCase{
		Case:                "DeleteCollection",
		Path:                "/" + testsCollection,
		ArrayRequestPayload: contents,
		ExpectedStatus:      http.StatusOK,
		ExpectedResponse: models.JsonResponse{
			Error:   false,
			Message: "Successfully deleted collection",
		},
	}

	chtc.RunContentDeleteTest(t)
}

func TestHandleDeleteClassFound(t *testing.T) {
	contentOne := models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "test-title1",
		Description: "test-description1",
		Body:        "test-body1",
		IsPublic:    false,
		Views:       10,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	contentTwo := models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class-wrong",
		Title:       "test-title2",
		Description: "test-description2",
		Body:        "test-body2",
		IsPublic:    true,
		Views:       0,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	contentThree := models.Content{
		Id:          uuid.New().String(),
		Class:       "test-class",
		Title:       "test-title1-diff",
		Description: "test-description1",
		Body:        "test-body1",
		IsPublic:    false,
		Views:       10,
		CreatorId:   uuid.New().String(),
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	var contents []models.Content
	contents = append(contents, contentOne, contentTwo, contentThree)

	chtc := ContentHandlerTestCase{
		Case:                "DeleteClass",
		Path:                "/" + testsCollection + "/class/test-class",
		ArrayRequestPayload: contents,
		ExpectedStatus:      http.StatusOK,
		ExpectedResponse: models.JsonResponse{
			Error:   false,
			Message: "Successfully deleted class",
		},
	}

	chtc.RunContentDeleteTest(t)
}
