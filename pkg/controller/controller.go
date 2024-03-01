package controller

import (
	"github.com/SvetlanaGrin/AvitoTest/pkg/service"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	services *service.Service
}

func NewControllers(services *service.Service) *Controllers {
	return &Controllers{services: services}
}

func (h *Controllers) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		balance := api.Group("/balance")
		{
			balance.GET("/:id", h.GetById)
		}
	}
	return router
}
