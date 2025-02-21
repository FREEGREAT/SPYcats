package handler

import (
	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:""`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, error{message})
}
