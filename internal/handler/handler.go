package handler

import (
	middleware "spy-cats/internal/middlewae"
	service "spy-cats/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger   middleware.Logger
}

func NewHandler(services *service.Service, log middleware.Logger) *Handler {
	return &Handler{services: services, logger: log}
}

func (h *Handler) InitRoutes(router *gin.Engine) *gin.Engine {
	cat := router.Group("/cat")
	{
		cat.POST("/", h.createCat)
		cat.DELETE("/:id", h.deleteCat)
		cat.GET("/:id", h.getCatById)
		cat.GET("/list", h.getListOfCats)
		cat.PUT("/", h.updateCatSalary)
	}
	mission := router.Group("/mission")
	{
		mission.POST("/", h.createMission)
		mission.DELETE("/:id", h.deleteMission)
		mission.GET("/:id", h.getMissionById)
		mission.GET("/list", h.getListOfMission)
		mission.PATCH("/complete/:id", h.completeMission)
	}
	target := router.Group("/target")
	{
		target.POST("/", h.createTarget)
		target.DELETE("/:id", h.deleteTarget)
		target.GET("/:id", h.getTargetById)
		target.GET("/by-mission/:id", h.getTargetListByMission)
		target.PATCH("/completed/:id", h.setTargetCompleted)
	}

	note := router.Group("/note")
	{
		note.POST("/", h.createNote)
		note.DELETE("/:id", h.deleteNote)
		note.PATCH("/", h.updateNote)
		note.GET("/by-target/:id", h.getListNoteByTarget)
		note.GET("/:id", h.getNoteById)

	}

	return router
}
