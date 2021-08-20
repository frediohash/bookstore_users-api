package app

import "github.com/frediohash/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.FindUser)
	router.GET("/user/:id", controllers.GetUser)
}
