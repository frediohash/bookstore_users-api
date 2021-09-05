package app

import (
	"github.com/frediohash/bookstore_users-api/controllers/ping"
	"github.com/frediohash/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/user", users.CreateUser0)
	router.GET("/user", users.FindUser)
	router.GET("/user/:user_id", users.GetUser)
	router.PUT("/user/:user_id", users.UpdateUser)
}
