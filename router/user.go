package router

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(c gin.RouterGroup) {
	c.POST("/login")
}
