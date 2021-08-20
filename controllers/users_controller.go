package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "Create User")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Get User")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusBadGateway, "Create User")
}
