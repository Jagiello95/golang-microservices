package services

import (
	"github.com/jagiello95/golang-microservices/mvc/domain"
	"github.com/jagiello95/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
