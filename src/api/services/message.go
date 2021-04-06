package services

import (
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/discord"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/messages"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/errors"
	discord_provider "github.com/yuricampolongo/microservices-rest-golang/src/api/providers/discord"
)

type messageService struct{}

type messageServiceInterface interface {
	SendMessage(input messages.MessageRequest) (*messages.MessageResponse, errors.ApiError)
}

var (
	MessageService messageServiceInterface
)

func init() {
	MessageService = &messageService{}
}

func (s *messageService) SendMessage(input messages.MessageRequest) (*messages.MessageResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := discord.Message{
		Content: input.Content,
		Tts:     false,
	}

	response, err := discord_provider.SendMessage(request)
	if err != nil {
		return nil, errors.NewApiError(err.Code, err.Message)
	}

	result := messages.MessageResponse{
		Code: response.Code,
	}
	return &result, nil
}
