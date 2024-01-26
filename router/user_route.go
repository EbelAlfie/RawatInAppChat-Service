package router

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(group *gin.RouterGroup) {
	group.POST("/login")
}
