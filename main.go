package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/obskur123/crent/src/config"
	"github.com/obskur123/crent/src/handlers"
	"github.com/obskur123/crent/src/helper"
	"github.com/obskur123/crent/src/repositories"
	"github.com/obskur123/crent/src/routes"
	"github.com/obskur123/crent/src/services"
)

func main() {

	err := godotenv.Load(".env")
	helper.ErrorPanic(err)

	color.Green("ᕕ( ᐛ )ᕗ gah mf damn dawg! starting the server real quick...")
	r := gin.Default()
	db := config.CreateConnection()
	
	userRepository := repositories.NewUserRepository(db)
	torrentRepository := repositories.NewTorrentRepository(db)

	userService := services.NewUserService(userRepository)
	torrentService := services.NewTorrentService(torrentRepository)

	routes.UserRoutes(r, userService)
	routes.TorrentRoutes(r, torrentService)

	r.LoadHTMLGlob("src/templates/*")
	r.GET("/ping", handlers.Ping)
	r.GET("/index", handlers.Home)
	r.Run()
}
