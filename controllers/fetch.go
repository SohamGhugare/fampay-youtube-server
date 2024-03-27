package controllers

import (
	"net/http"

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

	// TODO: fetch videos from youtube API

	c.JSON(http.StatusOK, gin.H{
		"query": query,
	})

}
