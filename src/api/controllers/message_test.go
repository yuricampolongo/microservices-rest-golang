package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/api_errors"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/clients/rest"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/messages"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/test"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockups()
	os.Exit(m.Run())
}

func TestSendMessageInvalidBody(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test.GetMockedContext(request, response)

	SendMessages(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := api_errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestSendMessageApiError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Invalid request","code": 50000}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"content":"message testing"}`))
	response := httptest.NewRecorder()
	c := test.GetMockedContext(request, response)

	SendMessages(c)

	assert.EqualValues(t, 50000, response.Code)
	apiErr, err := api_errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, 50000, apiErr.Status())
	assert.EqualValues(t, "Invalid request", apiErr.Message())
}

func TestSendMessage(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"content":"message testing"}`))
	response := httptest.NewRecorder()
	c := test.GetMockedContext(request, response)

	SendMessages(c)

	var result messages.MessageResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNoContent, result.Code)
}
