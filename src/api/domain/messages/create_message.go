package messages

import (
	"strings"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/api_errors"
)

type MessageRequest struct {
	Content string `json:"content"`
}

func (r *MessageRequest) Validate() api_errors.ApiError {
	r.Content = strings.TrimSpace(r.Content)
	if r.Content == "" {
		return api_errors.NewBadRequestError("invalid message content")
	}
	return nil
}

type MessageResponse struct {
	Content string `json:"content"`
	Code    int    `json:"code"`
}

type MessageSendResult struct {
	Response *MessageResponse    `json:"response"`
	Error    api_errors.ApiError `json:"error"`
}
