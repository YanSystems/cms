package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "github.com/YanSystems/cms/pkg/utils"
)

type ContentHandlerTestCase struct {
	Name             string
	RequestPayload   any
	ExpectedStatus   int
	ExpectedResponse utils.JsonResponse
}

func (chtc *ContentHandlerTestCase) RunContentHandlerTest(t *testing.T) {

	client, err := utils.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	t.Run(chtc.Name, func(t *testing.T) {
		requestPayload, err := json.Marshal(chtc.RequestPayload)
		if err != nil {
			t.Fatal("could not marshal request payload: %v", err)
		}

		server := httptest.NewServer(http.HandlerFunc(Handle))
	})
}
