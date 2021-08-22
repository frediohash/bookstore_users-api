package services

import (
	"github.com/frediohash/bookstore_users-api/domain/users"
	"github.com/frediohash/bookstore_users-api/utils/errors"
)

// //chapter1
// func CreateUser(user users.User) (*users.User, *errors.RestErr) {
// 	return &user, nil
// }

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
