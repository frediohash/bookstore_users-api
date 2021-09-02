package users

import (
	"fmt"
	"time"

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
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
	}

	now := time.Now()
	user.DateCreated = now.Format("2006-01-02T15:04:05Z")

	usersDB[user.Id] = user
	return nil
}
