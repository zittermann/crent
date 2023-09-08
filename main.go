package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/obskur123/crent/src/handlers"
)

func main() {
	color.Green("ᕕ( ᐛ )ᕗ gah mf damn dawg! starting the server real quick...")
	r := gin.Default()
	r.LoadHTMLGlob("src/templates/*")
	r.GET("/ping", handlers.Ping)
	r.GET("/index", handlers.Home)
	r.Run()
}
