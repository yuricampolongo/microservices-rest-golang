package app

import (
	"github.com/yuricampolongo/microservices-rest-golang/src/api/controllers"
)

func mapUrls() {
	router.POST("/messages", controllers.SendMessages)
}
