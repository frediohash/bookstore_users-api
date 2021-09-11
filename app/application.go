package app

import (
	"github.com/frediohash/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats-server/v2/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Log.Info("about to start application")
	router.Run(":8080")
}
