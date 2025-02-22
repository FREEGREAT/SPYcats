package handler

import (
	"net/http"
	"spy-cats/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createNote(c *gin.Context) {
	var input models.NoteModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.NoteRepository.CreateNote(c, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Note id": id,
	})
}

func (h *Handler) deleteNote(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	err = h.services.NoteRepository.DeleteNote(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getNoteById(c *gin.Context) {
	idStr := c.Param("id")
	var noteModel *models.NoteModel
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	noteModel, err = h.services.NoteRepository.GetNote(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, noteModel)
}

func (h *Handler) getListNoteByTarget(c *gin.Context) {
	idStr := c.Param("id")
	var noteModel []models.NoteModel
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}
	noteModel, err = h.services.NoteRepository.ListNotesByTarget(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, noteModel)
}

func (h *Handler) updateNote(c *gin.Context) {
	var input models.NoteModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.NoteRepository.UpdateNote(c, int64(input.ID), input.Content)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
