package domain

import "github.com/gin-gonic/gin"

type UserController interface {
	Login(c *gin.Context, param gin.Params)
}

type UserRepository interface {
	Login()
}
