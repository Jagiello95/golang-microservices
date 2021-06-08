package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jagiello95/golang-microservices/mvc/services"
	"github.com/jagiello95/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "userId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiError)

		resp.WriteHeader(apiError.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiError := services.GetUser(userId)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(jsonValue))
		//Handle the error and return to the client
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
