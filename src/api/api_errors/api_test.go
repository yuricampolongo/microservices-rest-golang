package api_errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApiError(t *testing.T) {
	apiErr := NewApiError(200, "message")

	assert.NotNil(t, apiErr)
	assert.EqualValues(t, 200, apiErr.Status())
	assert.EqualValues(t, "message", apiErr.Message())
}

func TestNewApiErrorFromBytes(t *testing.T) {
	apiErr, err := NewApiErrFromBytes([]byte(`{"status":200, "message":"message"}`))

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, 200, apiErr.Status())
	assert.EqualValues(t, "message", apiErr.Message())
}

func TestNewApiErrorInvalidBytes(t *testing.T) {
	apiErr, err := NewApiErrFromBytes([]byte(``))

	assert.NotNil(t, err)
	assert.Nil(t, apiErr)
	assert.EqualValues(t, "invalid json body", err.Error())
}

func TestNewInternaServerError(t *testing.T) {
	err := NewInternalServerError("message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "message", err.Message())
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "message", err.Message())
	assert.EqualValues(t, http.StatusNotFound, err.Status())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "message", err.Message())
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
}
