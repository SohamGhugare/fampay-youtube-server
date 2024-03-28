package main

import (
	"net/http"

	"github.com/SohamGhugare/fampay-youtube-server/controllers"
	"github.com/SohamGhugare/fampay-youtube-server/database"
	"github.com/SohamGhugare/fampay-youtube-server/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.SetupDatabaseClient()
}

func setupRoutes(r *gin.Engine) {
	r.GET("/api/v1/fetch", controllers.FetchVideos)
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

	// setting up routes
	setupRoutes(r)

	// running the server
	r.Run()

}
