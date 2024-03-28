package controllers

import (
	"net/http"

	"github.com/SohamGhugare/fampay-youtube-server/services"
	"github.com/gin-gonic/gin"
)

func FetchVideos(c *gin.Context) {
	// get query param
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query param is required",
		})
		return
	}

	videos := services.FetchSvc(query, 10)

	// for id, title := range videos {
	// 	log.Printf("[%v] %v\n", id, title)
	// }

	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"results": videos,
	})

}
