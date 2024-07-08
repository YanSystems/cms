package apitests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YanSystems/cms/pkg/models"
	"github.com/YanSystems/cms/pkg/server"
	utils "github.com/YanSystems/cms/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func (chtc *ContentHandlerTestCase) RunContentPostTest(t *testing.T) {
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

	r := api.NewRouter()

	payload, err := json.Marshal(chtc.RequestPayload)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", baseUrl+chtc.Path, bytes.NewBuffer(payload))
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

// func TestHandleCreateContentValid(t *testing.T) {
// 	content := models.Content{
// 		Class:       "test-class",
// 		Title:       "test-title",
// 		Description: "test-description",
// 		Body:        "test-body",
// 		IsPublic:    true,
// 		Views:       0,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	chtc := ContentHandlerTestCase{
// 		Path:           "/" + testsCollection,
// 		RequestPayload: content,
// 		ExpectedStatus: http.StatusCreated,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   false,
// 			Message: "Successfully created content",
// 		},
// 	}

// 	chtc.RunContentPostTest(t)
// }

// func TestHandleCreateContentIncorrectFields(t *testing.T) {
// 	content := models.Content{
// 		Class:       "test-class",
// 		Title:       "test-title",
// 		Description: "test-description",
// 		Body:        "test-body",
// 		IsPublic:    true,
// 		Views:       0,
// 		CreatorId:   uuid.New().String(),
// 	}

// 	chtc := ContentHandlerTestCase{
// 		Path:           "/" + testsCollection,
// 		RequestPayload: content,
// 		ExpectedStatus: http.StatusBadRequest,
// 		ExpectedResponse: models.JsonResponse{
// 			Error:   true,
// 			Message: "missing fields in request payload",
// 		},
// 	}

// 	t.Run("Missing class field", func(t *testing.T) {
// 		chtc.RequestPayload.Class = ""
// 		chtc.RunContentPostTest(t)
// 		chtc.RequestPayload.Class = "test-class"
// 	})

// 	t.Run("Missing title field", func(t *testing.T) {
// 		chtc.RequestPayload.Title = ""
// 		chtc.RunContentPostTest(t)
// 		chtc.RequestPayload.Title = "test-title"
// 	})

// 	t.Run("Missing description field", func(t *testing.T) {
// 		chtc.RequestPayload.Description = ""
// 		chtc.RunContentPostTest(t)
// 		chtc.RequestPayload.Description = "test-description"
// 	})

// 	t.Run("Missing body field", func(t *testing.T) {
// 		chtc.RequestPayload.Body = ""
// 		chtc.RunContentPostTest(t)
// 		chtc.RequestPayload.Body = "test-body"
// 	})

// 	t.Run("Incorrect views field", func(t *testing.T) {
// 		chtc.RequestPayload.Views = -1
// 		chtc.RunContentPostTest(t)
// 		chtc.RequestPayload.Views = 0
// 	})

// 	t.Run("Missing creator_id", func(t *testing.T) {
// 		chtc.RequestPayload.CreatorId = ""
// 		chtc.RunContentPostTest(t)
// 	})
// }
