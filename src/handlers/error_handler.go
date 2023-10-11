package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/obskur123/crent/src/utils"
)

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message string `json:"message"`
	Error string `json:"error"`
}

func NotFound(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusNotFound, &ErrorResponse {
		Timestamp: utils.GetDateTime(),
		Message: msg,
		Error: http.StatusText(http.StatusNotFound),
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse {
		Timestamp: utils.GetDateTime(),
		Message: msg,
		Error: http.StatusText(http.StatusBadRequest),
	})
}
