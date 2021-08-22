package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/frediohash/bookstore_users-api/domain/users"
	"github.com/frediohash/bookstore_users-api/services"
	"github.com/frediohash/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

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
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Get User")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusBadGateway, "Create User")
}
