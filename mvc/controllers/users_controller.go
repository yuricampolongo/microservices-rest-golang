package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yuricampolongo/microservices-rest-golang/mvc/services"
	"github.com/yuricampolongo/microservices-rest-golang/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		resp.WriteHeader(apiErr.Status)
		jsonValue, _ := json.Marshal(apiErr)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
