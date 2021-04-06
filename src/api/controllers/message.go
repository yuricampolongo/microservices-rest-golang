package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/api_errors"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/messages"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/services"
)

func SendMessage(c *gin.Context) {
	var request messages.MessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := api_errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.MessageService.SendMessage(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
