package apitests

import (
	"github.com/YanSystems/cms/pkg/models"
)

type ContentHandlerTestCase struct {
	Case              string
	Path              string
	ClassName         string
	AnyRequestPayload any
	RequestPayload    models.Content
	ExpectedStatus    int
	ExpectedResponse  models.JsonResponse
}

const baseUrl = "/contents"

var testsCollection = "services-tests"
