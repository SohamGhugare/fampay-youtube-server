package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// initializing gin server
	r := gin.Default()

	// test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	// running the server
	r.Run()

}
