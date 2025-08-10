package main

import (
	"github.com/debg48/gin_crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVaraibles()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
