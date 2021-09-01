package services

import (
	"github.com/frediohash/bookstore_users-api/domain/users"
	"github.com/frediohash/bookstore_users-api/utils/errors"
)

// //chapter1
// func CreateUser(user users.User) (*users.User, *errors.RestErr) {
// 	return &user, nil
// }

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
