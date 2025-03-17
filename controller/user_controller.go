package controller

import (
	"chat_service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repository domain.UserRepository
}

func NewUserController(repo domain.UserRepository) domain.UserController {
	return UserController{
		repository: repo,
	}
}

func (controller UserController) Login(c *gin.Context, param gin.Params) {
	_, usernameStatus := param.Get("username")
	_, passwordStatus := param.Get("pass")

	if !usernameStatus {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "username must not be null",
		})
	}
	if !passwordStatus {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "password must not be null",
		})
	}
}
