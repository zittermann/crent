package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/obskur123/crent/src/handlers"
	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/services"
)

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) FindByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	user, err := controller.service.FindByID(uint(id))

	if err != nil {
		handlers.NotFound(c, err.Error())
	}

	c.JSON(http.StatusOK, user)

}

func (controller *UserController) FindByName(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	name := c.Query("name")
	p, err := controller.service.FindByName(name, page, limit)
	
	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	c.JSON(http.StatusOK, p)

}

func (controller *UserController) FindByNickname(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))

	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	nickname := c.Query("nick")
	p, err := controller.service.FindByNickname(nickname, page, limit)
	
	if err != nil {
		handlers.BadRequest(c, err.Error())
	}

	c.JSON(http.StatusOK, p)

}

func (controller *UserController) Save(c *gin.Context) {

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		handlers.BadRequest(c, err.Error())
	}
	
	controller.service.Save(&user)

	c.JSON(http.StatusOK, user)

}
