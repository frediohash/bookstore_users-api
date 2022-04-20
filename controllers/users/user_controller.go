package users

import (
	"encoding/json"
	"fmt"
	"github.com/frediohash/bookstore_users-api/domain/users"
	"github.com/frediohash/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO Handler Error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO Handler Json Error
		return
	}
	//fmt.Println(err)
	//fmt.Println(string(bytes))
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		//TODO User Creation Error
		return
	}
	c.JSON(http.StatusCreated, result)
	fmt.Println(user)
	c.String(http.StatusNotImplemented, "implement me")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
