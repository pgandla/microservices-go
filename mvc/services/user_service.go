package services

import (
	"github.com/pgandla/microservices-go/mvc/domain"
	"github.com/pgandla/microservices-go/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppErr) {
	return domain.GetUser(userId)
}
