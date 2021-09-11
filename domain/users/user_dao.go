package users

import (
	"fmt"

	"github.com/frediohash/bookstore_users-api/datasources/mysql/users_db"
	"github.com/frediohash/bookstore_users-api/logger"
	"github.com/frediohash/bookstore_users-api/utils/date_utils"
	"github.com/frediohash/bookstore_users-api/utils/errors"
)

const (
	errorNoRows           = "no rows in result set"
	indexUniqueEmail      = "unique_Email"
	queryInsertUser       = ("INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);")
	queryGetUser          = ("SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;")
	queryUpdateUser       = ("UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;")
	queryDeleteUser       = ("DELETE FROM users WHERE id=?;")
	queryFindUserByStatus = ("SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;")
)

// get all dari database
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
		// return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
		//3
		// return mysql_utils.ParseError(getErr)
		//1
		// sqlErr, ok := getErr.(*mysql.MySQLError)
		// if !ok {
		// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying get user: %s", getErr.Error()))
		// }
		// fmt.Println(sqlErr)

		//2
		// if strings.Contains(err.Error(), errorNoRows) {
		// 	return errors.NewNotFoundError(
		// 		fmt.Sprintf("user %d not found", user.Id))
		// }
		// fmt.Println(err)
		// return errors.NewInternalServerError(
		// 	fmt.Sprintf("error when trying to get user %d: %s", user.Id, getErr.Error()))
	}
	return nil
}

// save ke database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
		// return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowDBFormat()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
		//3
		//because it was define in mysql_utils.go
		// return mysql_utils.ParseError(saveErr)

		//2
		// sqlErr, ok := saveErr.(*mysql.MySQLError)
		// if !ok {
		// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying save user: %s", err.Error()))
		// }
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
		// switch sqlErr.Number {
		// case 1062:
		// 	return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))

		//1
		// if strings.Contains(err.Error(), "email_UNIQUE") {
		// 	return errors.NewBadRequestError("email %s already exist")
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id when creating a new user", err)
		return errors.NewInternalServerError("database error")
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
		// return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError("database error")
		// return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
		// return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.NewInternalServerError("database error")
		// return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindUserByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user statement", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to prepare find user by status", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	//we want to know how many users in database
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
			// return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil

}

// make map access database
// var (
// 	usersDB = make(map[int64]*User)
// )

// first GET without database
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
