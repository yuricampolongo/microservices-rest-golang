package services

import (
	"net/http"
	"sync"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/api_errors"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/discord"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/messages"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/providers/discord_provider"
)

type messageService struct{}

type messageServiceInterface interface {
	Send(messages []messages.MessageRequest) (*[]messages.MessageResponse, api_errors.ApiError)
}

var (
	Message messageServiceInterface
)

func init() {
	Message = &messageService{}
}

func (s *messageService) Send(msgs []messages.MessageRequest) (*[]messages.MessageResponse, api_errors.ApiError) {
	input := make(chan messages.MessageSendResult)
	output := make(chan []messages.MessageResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleMessageSendResult(&wg, input, output)

	for _, m := range msgs {
		wg.Add(1)
		go s.handleMessageSend(m, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	return &result, nil

}

func (s *messageService) handleMessageSendResult(wg *sync.WaitGroup, input chan messages.MessageSendResult, output chan []messages.MessageResponse) {
	var results []messages.MessageResponse
	for ev := range input {
		if ev.Error != nil {
			response := messages.MessageResponse{
				Code: http.StatusBadRequest,
			}
			results = append(results, response)
			wg.Done()
			continue
		}

		response := messages.MessageResponse{
			Content: ev.Response.Content,
			Code:    ev.Response.Code,
		}
		results = append(results, response)
		wg.Done()
	}
	output <- results
}

func (s *messageService) handleMessageSend(input messages.MessageRequest, output chan messages.MessageSendResult) {
	if err := input.Validate(); err != nil {
		output <- messages.MessageSendResult{Error: err}
		return
	}

	response, err := discord_provider.SendMessage(discord.Message{
		Content: input.Content,
		Tts:     false,
	})

	if err != nil {
		output <- messages.MessageSendResult{Error: api_errors.NewApiError(err.Code, err.Message)}
		return
	}

	output <- messages.MessageSendResult{
		Response: &messages.MessageResponse{
			Code:    response.Code,
			Content: input.Content,
		},
	}
}
