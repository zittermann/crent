package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong boludito",
	})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "wtf is up bitch boy",
	})
}
