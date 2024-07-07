package apitests

import (
	"github.com/YanSystems/cms/pkg/models"
)

type ContentHandlerTestCase struct {
	CollectionName   string
	ClassName        string
	RequestPayload   models.Content
	ExpectedStatus   int
	ExpectedResponse models.JsonResponse
}
