package handler

import (
	"net/http"
	"spy-cats/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTarget(c *gin.Context) {
	var input models.TargetModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TargetRepository.CreateTarget(c, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Target id:": id,
	})
}
func (h *Handler) deleteTarget(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	err = h.services.TargetRepository.DeleteTarget(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getTargetById(c *gin.Context) {
	var targetModel *models.TargetModel
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	targetModel, err = h.services.TargetRepository.GetTarget(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, targetModel)
}

func (h *Handler) getTargetListByMission(c *gin.Context) {
	var targetModel []models.TargetModel
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	targetModel, err = h.services.TargetRepository.ListTargetsByMission(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, targetModel)
}

func (h *Handler) setTargetCompleted(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	err = h.services.TargetRepository.CompleteTarget(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
