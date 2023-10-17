package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/obskur123/crent/src/controllers"
	"github.com/obskur123/crent/src/services"
)

func TorrentRoutes(r *gin.Engine, service services.ITorrentService) {
	c := controllers.NewTorrentController(service)

	r.GET("/v1/torrents/:id", c.FindByID)
	r.GET("/v1/torrents/title", c.FindByTitle)
	r.POST("/v1/torrents", c.Save)

}