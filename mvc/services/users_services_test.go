package services

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/domain"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNoUserFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message: "user 0 not found",
			Status:  http.StatusNotFound,
			Code:    "not_found",
		}
	}

	user, err := UserServices.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestGetUser(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Geralt",
			LastName:  "Of Rivia",
			Email:     "geralt@witcher.com",
		}, nil
	}

	user, err := UserServices.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Geralt", user.FirstName)
	assert.EqualValues(t, "Of Rivia", user.LastName)
	assert.EqualValues(t, "geralt@witcher.com", user.Email)
}
