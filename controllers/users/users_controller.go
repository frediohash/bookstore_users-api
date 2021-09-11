package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/frediohash/bookstore_users-api/domain/users"
	"github.com/frediohash/bookstore_users-api/services"
	"github.com/frediohash/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("userId should be a number")
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO something here
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// fmt.Println(err.Error())
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		// restErr := errors.NewBadRequestError("invalid json body"){
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-Public") == "true"))
}

func CreateUser0(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-Public") == "true"))
}

func GetUser(c *gin.Context) {
	// not solid principal, take from getUserId
	// userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// if userErr != nil {
	// 	err := errors.NewBadRequestError("invalid user id")
	// 	c.JSON(err.Status, err)
	// 	return
	// }

	// solid principal, take from getUserId
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-Public") == "true"))
}

func FindUser(c *gin.Context) {
	c.String(http.StatusBadGateway, "Find User")
}

func UpdateUser(c *gin.Context) {
	// not solid principal, take id param
	// userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// if userErr != nil {
	// 	err := errors.NewBadRequestError("invalid user id")
	// 	c.JSON(err.Status, err)
	// 	return
	// }

	// solid principal, take from getUserId
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	//get json and return json
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	// solid principal, take from getUserId
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-Public") == "true"))
}
