package app

import (
	"github.com/frediohash/bookstore_users-api/controllers/ping"
	"github.com/frediohash/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/user", users.CreateUser)
	router.GET("/user", users.FindUser)
	router.GET("/user/:id", users.GetUser)
}
