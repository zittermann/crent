package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/handlers"
	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/services"
)

type TorrentController struct {
	service services.ITorrentService
}

func NewTorrentController(
	service services.ITorrentService) *TorrentController {
		return &TorrentController {
			service: service,
		}
}

func (contoller *TorrentController) FindByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	torrent, err := contoller.service.FindByID(uint(id))

	if err != nil { 
		handlers.NotFound(c, err.Error()) 
	}

	c.JSON(http.StatusOK, torrent)

}

func (controller *TorrentController) FindByTitle(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	p := response.Page{ Page: page, Limit: limit }

	title := c.Query("title")
	err = controller.service.FindByTitle(title, &p)
	
	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	c.JSON(http.StatusOK, p)

}

func (controller *TorrentController) Save(c *gin.Context) {

	torrent := models.Torrent{}

	if err := c.ShouldBindJSON(&torrent); err != nil {
		handlers.BadRequest(c, err.Error())
	}

	controller.service.Save(&torrent)

	c.JSON(http.StatusCreated, torrent)

}
