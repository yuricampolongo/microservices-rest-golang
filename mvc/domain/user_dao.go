package domain

import (
	"fmt"
	"net/http"

	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Geralt", LastName: "Of Rivia", Email: "geralt@witcher.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
