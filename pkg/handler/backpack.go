package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strimenko/store-strimenko/models"
)

func (h *Handler) createBackpack(c *gin.Context) {
	var input models.Backpack

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.Backpack.CreateBackpack(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"itemid": itemId,
	})
}

type getAllBackpackResponse struct {
	Data []models.Backpack
}

func (h *Handler) getAllBackpacks(c *gin.Context) {
	lists, err := h.services.Backpack.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBackpackResponse{
		Data: lists,
	})
}

func (h *Handler) getBackpackById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("itemid"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid itemid param")
		return
	}

	lists, err := h.services.Backpack.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (h *Handler) updateBackpack(c *gin.Context) {

}

func (h *Handler) deleteBackpack(c *gin.Context) {

}
