package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginController(c *gin.Context, param gin.Params) {
	username, usernameStatus := param.Get("username")
	password, passwordStatus := param.Get("pass")

	if !usernameStatus {
		c.JSON(http.StatusNotFound, gin.H{
			{"message": "username must not be null"}
		})
	}
	if !passwordStatus {
		c.JSON(http.StatusNotFound, gin.H{
			{"message": "password must not be null"}
		})
	}
}
