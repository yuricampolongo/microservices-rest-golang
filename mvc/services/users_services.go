package services

import (
	"github.com/yuricampolongo/microservices-rest-golang/mvc/domain"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

type userServices struct{}

var (
	UserServices userServices
)

func (t *userServices) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
