package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with ID 0")
	assert.NotNil(t, err, "we were expecting an error when user ID is 0")
	assert.EqualValues(t, err.Status, http.StatusNotFound)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestGetUser(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Geralt", user.FirstName)
	assert.EqualValues(t, "Of Rivia", user.LastName)
	assert.EqualValues(t, "geralt@witcher.com", user.Email)
}
