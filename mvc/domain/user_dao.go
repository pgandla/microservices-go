package domain

import (
	"github.com/pgandla/microservices-go/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{
			ID:        123,
			FirstName: "Pradeep",
			LastName:  "Gandla",
			Email:     "pradeepdoit@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.AppErr)  {
	user := users[userId]
	if user == nil {
		return nil, &utils.AppErr{
			Message:    "user not found",
			StatusCode: 404,
			ErrCode:    "not_found",
		}
	}
	return user, nil
}
