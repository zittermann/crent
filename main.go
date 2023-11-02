package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"

	"github.com/obskur123/crent/src/config"
	"github.com/obskur123/crent/src/handlers"
	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/repositories"
	"github.com/obskur123/crent/src/routes"
	"github.com/obskur123/crent/src/services"
)

func main() {

	color.Green("ᕕ( ᐛ )ᕗ gah mf damn dawg! starting the server real quick...")
	r := gin.Default()
	db := config.CreateConnection()

	db.Table("users").AutoMigrate(&models.User{})
	db.Table("torrents").AutoMigrate(&models.Torrent{})
	
	userRepository := repositories.NewUserRepository(db)
	torrentRepository := repositories.NewTorrentRepository(db)

	userService := services.NewUserService(userRepository)
	torrentService := services.NewTorrentService(torrentRepository)

	routes.UserRoutes(r, userService)
	routes.TorrentRoutes(r, torrentService)

	// r.LoadHTMLGlob("src/templates/*")
	r.GET("/ping", handlers.Ping)
	r.GET("/index", handlers.Home)
	r.Run()
}
