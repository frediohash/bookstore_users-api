package users

import (
	"fmt"
	"github.com/frediohash/bookstore_users-api/domain/users"
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
		return
	}
	fmt.Println(err)
	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented, "implement me")

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
