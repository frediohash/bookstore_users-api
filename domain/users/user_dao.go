package users

import (
	"fmt"

	"github.com/frediohash/bookstore_users-api/utils/errors"
)

// buat map akses database
var (
	usersDB = make(map[int64]*User)
)

// untuk ke database
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.Id))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	return nil
}
