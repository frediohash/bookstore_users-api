package app

import (
	"github.com/frediohash/bookstore_users-api/controllers/ping"
	"github.com/frediohash/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:userid", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)

}
