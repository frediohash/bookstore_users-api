package users

import (
	"fmt"

	"github.com/frediohash/bookstore_users-api/datasources/mysql/users_db"
	"github.com/frediohash/bookstore_users-api/utils/errors"
	"github.com/frediohash/bookstore_users-api/utils/errors/date_utils"
)

const (
	queryInsertUser = ("INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);")
)

// buat map akses database
var (
	usersDB = make(map[int64]*User)
)

// untuk ke database
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to save user: %s", err.Error())
		)
	}
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
	}

	// now := time.Now()
	// user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
