package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/services"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		c.JSON(apiErr.Status, apiErr)
		return
	}

	user, apiErr := services.UserServices.GetUser(userId)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
