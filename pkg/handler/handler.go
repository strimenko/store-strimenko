package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/strimenko/store-strimenko/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		helmet := api.Group("/helmet")
		{
			helmet.POST("/", h.createHelmet)
			helmet.GET("/", h.getAllHelmets)
			helmet.GET("/:itemid", h.getHelmetById)
			helmet.PUT("/:itemid", h.updateHelmet)
			helmet.DELETE("/:itemid", h.deleteHelmet)
		}

		bodyarmor := api.Group("/bodyarmor")
		{
			bodyarmor.POST("/", h.createBodyarmor)
			bodyarmor.GET("/", h.getAllBodyarmors)
			bodyarmor.GET("/:itemid", h.getBodyarmorById)
			bodyarmor.PUT("/:itemid", h.updateBodyarmor)
			bodyarmor.DELETE("/:itemid", h.deleteBodyarmor)
		}

		backpack := api.Group("/backpack")
		{
			backpack.POST("/", h.createBackpack)
			backpack.GET("/", h.getAllBackpacks)
			backpack.GET("/:itemid", h.getBackpackById)
			backpack.PUT("/:itemid", h.updateBackpack)
			backpack.DELETE("/:itemid", h.deleteBackpack)
		}
	}
	return router
}
