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
		"sitename": "wtf is up bitch boy",
		"msg":      "this shit wack",
	})
}
