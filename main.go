package main

import (
	"net/http"

	"github.com/SohamGhugare/fampay-youtube-server/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

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
