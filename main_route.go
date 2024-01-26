package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.Group("/main")

	server.Run("8080")
}
