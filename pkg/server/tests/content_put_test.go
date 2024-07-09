package apitests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (chtc *ContentHandlerTestCase) RunContentPutTest(t *testing.T) {
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

	repo := repositories.ContentRepository{DB: api.DB}
	fmt.Println(chtc.ArrayRequestPayload[0])
	id, err := repo.CreateContent(testsCollection, &chtc.ArrayRequestPayload[0])
	assert.NoError(t, err)

	payload, err := json.Marshal(chtc.ArrayRequestPayload[1])
	assert.NoError(t, err)

	fmt.Println("ZE URL:", baseUrl+chtc.Path+"/id/"+id)
	req, err := http.NewRequest("PUT", baseUrl+chtc.Path+"/id/"+id, bytes.NewBuffer(payload))
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

func TestHandlerUpdateContentPartial(t *testing.T) {
	createPayload := models.Content{
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

	updatePayload := models.Content{
		Title: "updated-title",
	}

	var updateContents []models.Content
	updateContents = append(updateContents, createPayload)
	updateContents = append(updateContents, updatePayload)

	chtc := ContentHandlerTestCase{
		Path:                "/" + testsCollection,
		ArrayRequestPayload: updateContents,
		ExpectedStatus:      http.StatusOK,
		ExpectedResponse: models.JsonResponse{
			Error:   false,
			Message: "Successfully updated content",
		},
	}

	chtc.RunContentPutTest(t)
}
