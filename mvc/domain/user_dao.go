package domain

import (
	"fmt"
	"net/http"

	"github.com/jagiello95/golang-microservices/mvc/utils"
)

//mocked database
var (
	users = map[int64]*User{
		123: {
			Id:        1,
			FirstName: "Andrzej",
			LastName:  "Jagiello",
			Email:     "and.jagiello@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
