package services

import (
	"github.com/AhmedSafwat1/micro_service_go-api/domain/users"
	"github.com/AhmedSafwat1/micro_service_go-api/utils/errors"
	"strings"
)

func GetUser(userId int64) (*users.User, *errors.RestError)  {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return  nil , errors.NewBadRequest("id not vlaid")
	}
	return  &result , nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError)  {
	user.Email    = strings.TrimSpace(strings.ToLower(user.Email))
	if err := user.Validate(); err !=nil {
		return nil, err
	}
	if err := user.Save(); err !=nil {
		return nil, err
	}
	return &user, nil
}
