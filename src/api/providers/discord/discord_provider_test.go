package discord_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/clients/rest"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/discord"

	"github.com/stretchr/testify/assert"
)

const (
	sendMessageUrl = "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH"
)

func TestMain(m *testing.M) {
	rest.StartMockups()
	os.Exit(m.Run())
}

func TestSendMessageErrorClient(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        sendMessageUrl,
		HttpMethod: http.MethodPost,
		Err:        errors.New("error to send message to channel"),
	})

	response, err := SendMessage(discord.Message{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "error to send message to channel", err.Message)
}

func TestSendMessageErrorInvalidResponseBody(t *testing.T) {
	rest.FlushMockups()
	invalidCloser, _ := os.Open("1454-555")
	rest.AddMockup(rest.Mock{
		Url:        sendMessageUrl,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       invalidCloser,
		},
	})

	response, err := SendMessage(discord.Message{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid body", err.Message)
}

func TestSendMessageErrorInvalidErrorResponseBody(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        sendMessageUrl,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := SendMessage(discord.Message{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid json error response body", err.Message)
}

func TestSendMessageErrorInvalidRequest(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        sendMessageUrl,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Cannot send an empty message", "code": 50006}`)),
		},
	})

	response, err := SendMessage(discord.Message{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, 50006, err.Code)
	assert.EqualValues(t, "Cannot send an empty message", err.Message)
}

func TestSendMessage(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        sendMessageUrl,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       ioutil.NopCloser(strings.NewReader(``)),
		},
	})

	response, err := SendMessage(discord.Message{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusNoContent, response.Code)
}
