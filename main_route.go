package main

import (
	"chat_service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	publicRoute := server.Group("/main")
	{
		server.Group("/user", router.UserRoute(publicRoute))
	}

	server.Run("8080")
}
