package app

import (
	"github.com/yuricampolongo/microservices-rest-golang/mvc/controllers"
)

func mapUrls() {
	router.GET("/user/:user_id", controllers.GetUser)
}
