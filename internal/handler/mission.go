package handler

import (
	"net/http"
	"spy-cats/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createMission(c *gin.Context) {
	var input models.MissionModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.MissionRepository.CreateMission(c, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Mission id": id,
	})
}
func (h *Handler) deleteMission(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.services.MissionRepository.DeleteMission(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getMissionById(c *gin.Context) {
	idStr := c.Param("id")

	var mission *models.MissionModel

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	mission, err = h.services.MissionRepository.GetMission(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, mission)
}
func (h *Handler) getListOfMission(c *gin.Context) {
	var missions []models.MissionModel
	missions, err := h.services.MissionRepository.ListMissions(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, missions)
}
func (h *Handler) completeMission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	err = h.services.MissionRepository.CompleteMission(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)

}
