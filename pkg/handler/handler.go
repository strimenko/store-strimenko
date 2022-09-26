package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
			helmet.GET("/:iditem", h.getHelmetById)
			helmet.PUT("/:iditem", h.updateHelmet)
			helmet.DELETE("/:iditem", h.deleteHelmet)
		}

		bodyarmor := api.Group("/bodyarmor")
		{
			bodyarmor.POST("/", h.createBodyarmor)
			bodyarmor.GET("/", h.getAllBodyarmors)
			bodyarmor.GET("/:iditem", h.getBodyarmorById)
			bodyarmor.PUT("/:iditem", h.updateBodyarmor)
			bodyarmor.DELETE("/:iditem", h.deleteBodyarmor)
		}

		backpack := api.Group("/backpack")
		{
			backpack.POST("/", h.createBackpack)
			backpack.GET("/", h.getAllBackpacks)
			backpack.GET("/:iditem", h.getBackpackById)
			backpack.PUT("/:iditem", h.updateBackpack)
			backpack.DELETE("/:iditem", h.deleteBackpack)
		}
	}
	return router
}
