package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Geralt", LastName: "Of Rivia", Email: "geralt@witcher.com"},
	}

	UserDao userServiceInterface
)

func init() {
	UserDao = &userDao{}
}

type userServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("user_dao.go -> GetUser")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
