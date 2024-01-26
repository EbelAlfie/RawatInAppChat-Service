package main

import (
	"chat_service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Group("/main")
	{
		server.Group("user", router.UserRoute)
	}

	server.Run("8080")
}
