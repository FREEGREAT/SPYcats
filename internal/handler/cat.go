package handler

import (
	"net/http"
	"spy-cats/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createCat(c *gin.Context) {
	var input models.CatModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	breed, err := h.services.CatApi.IsValidBreed(input.Breed)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !breed {
		newErrorResponse(c, http.StatusBadRequest, "Breeds does not exits")
		return
	}
	id, err := h.services.CatRepository.CreateSpyCat(c, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Cat id":           id,
		"Breed validation": breed,
	})

}

func (h *Handler) deleteCat(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.services.CatRepository.DeleteSpyCat(c, &id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)

}
func (h *Handler) getCatById(c *gin.Context) {
	idStr := c.Param("id")
	var cat *models.CatModel

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	cat, err = h.services.CatRepository.GetSpyCat(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cat)

}
func (h *Handler) getListOfCats(c *gin.Context) {
	var cats []models.CatModel

	cats, err := h.services.CatRepository.ListSpyCats(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cats)

}
func (h *Handler) updateCatSalary(c *gin.Context) {
	var reqMod models.UpdateSalaryRequest

	if err := c.ShouldBindJSON(&reqMod); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	response, err := h.services.CatRepository.UpdateSpyCatSalary(c, reqMod.ID, reqMod.Salary)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"NewSalary": response})
}
