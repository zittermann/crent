package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/obskur123/crent/src/controllers"
	"github.com/obskur123/crent/src/services"
)

func UserRoutes(r *gin.Engine, service services.IUserService) {
	c := controllers.NewUserController(service)

	r.GET("api/v1/users/:id", c.FindByID)
	r.GET("api/v1/users/name", c.FindByName)
	r.GET("api/v1/users/nick", c.FindByNickname)
	r.POST("api/v1/users/", c.Save)
}