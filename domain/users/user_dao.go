package users

import (
	"fmt"
	"strings"

	"github.com/frediohash/bookstore_users-api/datasources/mysql/users_db"
	"github.com/frediohash/bookstore_users-api/utils/errors"
	"github.com/frediohash/bookstore_users-api/utils/errors/date_utils"
)

const (
	errorNoRows      = "no rows in result set"
	indexUniqueEmail = "unique_Email"
	queryInsertUser  = ("INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);")
	queryGetUser     = ("SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?")
)

// get all dari database
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
				return errors.NewNotFoundError(
					fmt.Sprintf("user %s not found", user.Id)
				)
		}
		fmt.Println(err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}
	return nil
}

// save ke database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError("email %s already exist")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

// buat map akses database
// var (
// 	usersDB = make(map[int64]*User)
// )

// get yang pertama tanpa database
// func (user *User) Get() *errors.RestErr {
// 	if err := users_db.Client.Ping(); err != nil {
// 		panic(err)
// 	}
// 	result := usersDB[user.Id]
// 	if result == nil {
// 		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
// 	}
// 	user.Id = result.Id
// 	user.FirstName = result.FirstName
// 	user.LastName = result.LastName
// 	user.Email = result.Email
// 	user.DateCreated = result.DateCreated

// 	return nil
// }

// save yang pertama tanpa database
// func (user *User) Save() *errors.RestErr {
// 	stmt, err := users_db.Client.Prepare(queryInsertUser)
// 	if err != nil {
// 		return errors.NewInternalServerError(err.Error())
// 	}
// 	defer stmt.Close()
// result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
// 	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
// 	if err != nil {
// 		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
// 	}
// 	userId, err := insertResult.LastInsertId()
// 	if err != nil {
// 		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
// 	}
// 	current := usersDB[user.Id]
// 	if current != nil {
// 		if current.Email == user.Email {
// 			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
// 		}
// 		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
// 	}

// 	// now := time.Now()
// 	// user.DateCreated = now.Format("2006-01-02T15:04:05Z")
// 	user.DateCreated = date_utils.GetNowString()

// 	usersDB[user.Id] = user
// 	return nil
// }
