package services

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/api_errors"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/clients/rest"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/messages"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockups()
	os.Exit(m.Run())
}

func TestSend(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		},
	})

	requests := []messages.MessageRequest{
		{Content: "message 1"},
		{Content: "message 2"},
	}

	result, err := Message.Send(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.EqualValues(t, http.StatusNoContent, (*result)[0].Code)
	assert.Regexp(t, "message \\d", (*result)[0].Content)

	assert.EqualValues(t, http.StatusNoContent, (*result)[1].Code)
	assert.Regexp(t, "message \\d", (*result)[1].Content)
}

func TestSendMessageWithError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Err:        errors.New("error to send message"),
	})

	requests := []messages.MessageRequest{
		{Content: "  "},
		{Content: ""},
	}

	result, err := Message.Send(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.EqualValues(t, http.StatusBadRequest, (*result)[0].Code)
	assert.Empty(t, (*result)[0].Content)

	assert.EqualValues(t, http.StatusBadRequest, (*result)[1].Code)
	assert.Empty(t, (*result)[1].Content)
}

func TestHandleMessageSendInvalidRequest(t *testing.T) {
	input := messages.MessageRequest{}
	service := messageService{}
	output := make(chan messages.MessageSendResult)
	go service.handleMessageSend(input, output)

	response := <-output

	assert.Nil(t, response.Response)
	assert.NotNil(t, response.Error)
	assert.EqualValues(t, "invalid message content", response.Error.Message())
	assert.EqualValues(t, http.StatusBadRequest, response.Error.Status())
}

func TestHandleMessageSendDiscordError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockup(rest.Mock{
		Url:        "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH",
		HttpMethod: http.MethodPost,
		Err:        errors.New("error to send message"),
	})

	request := messages.MessageRequest{
		Content: "message with error",
	}

	service := messageService{}
	output := make(chan messages.MessageSendResult)
	go service.handleMessageSend(request, output)

	response := <-output

	assert.Nil(t, response.Response)
	assert.NotNil(t, response.Error)
	assert.EqualValues(t, "error to send message", response.Error.Message())
	assert.EqualValues(t, http.StatusInternalServerError, response.Error.Status())
}

func TestHandleMessage(t *testing.T) {
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
		Content: "message ",
	}

	service := messageService{}
	output := make(chan messages.MessageSendResult)
	go service.handleMessageSend(request, output)

	response := <-output

	assert.NotNil(t, response.Response)
	assert.Nil(t, response.Error)
	assert.EqualValues(t, "message", response.Response.Content)
	assert.EqualValues(t, http.StatusNoContent, response.Response.Code)
}

func TestHandleMessageSendResult(t *testing.T) {
	service := messageService{}
	input := make(chan messages.MessageSendResult)
	output := make(chan []messages.MessageResponse)
	var wg sync.WaitGroup

	go service.handleMessageSendResult(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- messages.MessageSendResult{
			Response: &messages.MessageResponse{
				Content: "Message sent",
				Code:    http.StatusNoContent,
			},
		}
	}()

	wg.Wait()
	close(input)

	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusNoContent, result[0].Code)
	assert.EqualValues(t, "Message sent", result[0].Content)
}

func TestHandleMessageSendResultWithError(t *testing.T) {
	service := messageService{}
	input := make(chan messages.MessageSendResult)
	output := make(chan []messages.MessageResponse)
	var wg sync.WaitGroup

	go service.handleMessageSendResult(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- messages.MessageSendResult{
			Error: api_errors.NewBadRequestError("Invalid request"),
		}
	}()

	wg.Wait()
	close(input)

	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusBadRequest, result[0].Code)
}
