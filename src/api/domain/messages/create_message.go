package messages

import (
	"strings"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/errors"
)

type MessageRequest struct {
	Content string `json:"content"`
}

func (r *MessageRequest) Validate() errors.ApiError {
	r.Content = strings.TrimSpace(r.Content)
	if r.Content == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

type MessageResponse struct {
	Code int `json:"code"`
}
