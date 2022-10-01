package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strimenko/store-strimenko/models"
)

func (h *Handler) createHelmet(c *gin.Context) {
	var input models.Helmet

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.Helmet.CreateHelmet(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"itemid": itemId,
	})

}

type getAllHelmetResponse struct {
	Data []models.Helmet
}

func (h *Handler) getAllHelmets(c *gin.Context) {
	lists, err := h.services.Helmet.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllHelmetResponse{
		Data: lists,
	})
}

func (h *Handler) getHelmetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("itemid"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid itemid param")
		return
	}

	lists, err := h.services.Helmet.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (h *Handler) updateHelmet(c *gin.Context) {

}

func (h *Handler) deleteHelmet(c *gin.Context) {

}
