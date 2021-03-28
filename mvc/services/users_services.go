package services

import (
	"github.com/yuricampolongo/microservices-rest-golang/mvc/domain"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
