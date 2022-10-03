package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strimenko/store-strimenko/models"
)

func (h *Handler) createBodyarmor(c *gin.Context) {
	var input models.Bodyarmor

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.Bodyarmor.CreateBodyarmor(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"itemid": itemId,
	})

}

type getAllBodyarmorResponse struct {
	Data []models.Bodyarmor
}

func (h *Handler) getAllBodyarmors(c *gin.Context) {
	lists, err := h.services.Bodyarmor.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBodyarmorResponse{
		Data: lists,
	})
}

func (h *Handler) getBodyarmorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("itemid"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid itemid param")
		return
	}

	lists, err := h.services.Bodyarmor.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (h *Handler) updateBodyarmor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("itemid"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid itemid param")
		return
	}

	var input models.UpdateItem
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Bodyarmor.Update(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})
}

func (h *Handler) deleteBodyarmor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("itemid"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid itemid param")
		return
	}

	err = h.services.Bodyarmor.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
