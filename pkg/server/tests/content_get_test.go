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

func (chtc *ContentHandlerTestCase) RunContentGetTest(t *testing.T) {
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

	repo := repositories.ContentRepository{
		DB: api.DB,
	}

	r := api.NewRouter()

	switch chtc.Case {
	case "GetContentFound":
		chtc.RequestPayload.Id = uuid.New().String()
		chtc.RequestPayload.UpdatedAt = time.Now().UTC()
		chtc.RequestPayload.CreatedAt = time.Now().UTC()
		id, err := repo.CreateContent(testsCollection, &chtc.RequestPayload)
		assert.NoError(t, err)
		chtc.Path += "/" + id
	case "GetContentNotFound":
		chtc.Path += "/non-existent-content"
	case "GetCollectionFound", "GetClassFound":
		contents := chtc.ArrayRequestPayload
		for _, content := range contents {
			content.Id = uuid.New().String()
			content.UpdatedAt = time.Now().UTC()
			content.CreatedAt = time.Now().UTC()
			_, err := repo.CreateContent(testsCollection, &content)
			assert.NoError(t, err)
		}
	case "GetCollectionNotFound":
	default:
		t.Errorf("Invalid case: %v", chtc.Case)
	}

	req, err := http.NewRequest("GET", baseUrl+chtc.Path, nil)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var responsePayload models.JsonResponse
	_ = json.NewDecoder(rr.Body).Decode(&responsePayload)

	assert.Equal(t, chtc.ExpectedStatus, rr.Code, "HTTP status code mismatch")
	assert.Equal(t, chtc.ExpectedResponse.Error, responsePayload.Error, "Error field mismatch")
	assert.Equal(t, chtc.ExpectedResponse.Message, responsePayload.Message, "Message field mismatch")

	switch chtc.Case {
	case "GetContentFound":
		if !responsePayload.Error {
			if data, ok := responsePayload.Data.(map[string]interface{}); ok {
				assert.Equal(t, chtc.RequestPayload.Id, data["id"], "ID mismatch")
				assert.Equal(t, chtc.RequestPayload.Class, data["class"], "class mismatch")
				assert.Equal(t, chtc.RequestPayload.Title, data["title"], "title mismatch")
				assert.Equal(t, chtc.RequestPayload.Description, data["description"], "description mismatch")
				assert.Equal(t, chtc.RequestPayload.Body, data["body"], "body mismatch")
				assert.Equal(t, chtc.RequestPayload.IsPublic, data["is_public"], "is_public mismatch")
				assert.Equal(t, float64(chtc.RequestPayload.Views), data["views"], "views mismatch")
				assert.Equal(t, chtc.RequestPayload.CreatorId, data["creator_id"], "creator_id mismatch")
			} else {
				t.Error("Type assertion failed")
			}
		}
	}

	_, err = repo.DeleteContent(testsCollection, chtc.RequestPayload.Id)
	assert.NoError(t, err)
}

// func TestHandleGetContentFound(t *testing.T) {
// 	content := models.Content{
// 		Class:       "test-class",
// 		Title:       "test-title",
// 		Description: "test-description",
// 		Body:        "test-body",
// 		IsPublic:    true,
// 		Views:       32,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	chtc := ContentHandlerTestCase{
// 		Case:           "GetContentFound",
// 		Path:           "/" + testsCollection + "/id",
// 		RequestPayload: content,
// 		ExpectedStatus: http.StatusOK,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   false,
// 			Message: "Successfully retrieved content",
// 		},
// 	}

// 	chtc.RunContentGetTest(t)
// }

// func TestHandleGetContentNotFound(t *testing.T) {
// 	chtc := ContentHandlerTestCase{
// 		Case:           "GetContentNotFound",
// 		Path:           "/" + testsCollection + "/id",
// 		ExpectedStatus: http.StatusBadRequest,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   true,
// 			Message: "content not found",
// 		},
// 	}

// 	chtc.RunContentGetTest(t)
// }

// func TestHandleGetCollectionFound(t *testing.T) {
// 	var contents []models.Content

// 	contentOne := models.Content{
// 		Class:       "test-class1",
// 		Title:       "test-title1",
// 		Description: "test-description1",
// 		Body:        "test-body1",
// 		IsPublic:    true,
// 		Views:       16,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	contentTwo := models.Content{
// 		Class:       "test-class2",
// 		Title:       "test-title2",
// 		Description: "test-description2",
// 		Body:        "test-body2",
// 		IsPublic:    true,
// 		Views:       16,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	contents = append(contents, contentOne, contentTwo)

// 	chtc := ContentHandlerTestCase{
// 		Case:                "GetCollectionFound",
// 		Path:                "/" + testsCollection,
// 		ArrayRequestPayload: contents,
// 		ExpectedStatus:      http.StatusOK,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   false,
// 			Message: "Successfully retrieved collection",
// 		},
// 	}

// 	chtc.RunContentGetTest(t)
// }

// func TestHandleGetCollectionNotFound(t *testing.T) {
// 	chtc := ContentHandlerTestCase{
// 		Case:           "GetCollectionNotFound",
// 		Path:           "/non-existent-collection",
// 		ExpectedStatus: http.StatusOK,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   false,
// 			Message: "Successfully retrieved collection", // should still get an empty array
// 		},
// 	}

// 	chtc.RunContentGetTest(t)
// }

// func TestHandleGetClassFound(t *testing.T) {
// 	var contents []models.Content

// 	contentOne := models.Content{
// 		Class:       "test-class1",
// 		Title:       "test-title1",
// 		Description: "test-description1",
// 		Body:        "test-body1",
// 		IsPublic:    true,
// 		Views:       16,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	contentTwo := models.Content{
// 		Class:       "test-class2",
// 		Title:       "test-title2",
// 		Description: "test-description2",
// 		Body:        "test-body2",
// 		IsPublic:    true,
// 		Views:       16,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	contents = append(contents, contentOne, contentTwo)

// 	class := "the-test-class"
// 	chtc := ContentHandlerTestCase{
// 		Case:                "GetClassFound",
// 		Path:                "/" + testsCollection + "/class/" + class,
// 		ArrayRequestPayload: contents,
// 		ExpectedStatus:      http.StatusOK,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   false,
// 			Message: "Successfully retrieved class",
// 		},
// 	}

// 	chtc.RunContentGetTest(t)
// }

// func TestHandleGetClassNotFound(t *testing.T) {
// 	chtc := ContentHandlerTestCase{
// 		Case:           "GetCollectionNotFound",
// 		Path:           "/" + testsCollection + "/class/non-existent-class",
// 		ExpectedStatus: http.StatusOK,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   false,
// 			Message: "Successfully retrieved class", // should still get an empty array
// 		},
// 	}

// 	chtc.RunContentGetTest(t)
// }
