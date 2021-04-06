package services

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/clients/rest"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/messages"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockups()
	os.Exit(m.Run())
}

func TestSendMessageNoContent(t *testing.T) {
	request := messages.MessageRequest{}

	response, err := MessageService.SendMessage(request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid message content", err.Message())
}

func TestSendMessageApiError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Err:        errors.New("error to send message"),
	})

	request := messages.MessageRequest{
		Content: "message with error",
	}

	response, err := MessageService.SendMessage(request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error to send message", err.Message())
}

func TestSendMessage(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		},
	})

	request := messages.MessageRequest{
		Content: "message mock",
	}

	response, err := MessageService.SendMessage(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusNoContent, response.Code)
}
