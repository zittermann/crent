package main

import (
	"github.com/gin-gonic/gin"
	"github.com/obskur123/crent/src/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run()
}
