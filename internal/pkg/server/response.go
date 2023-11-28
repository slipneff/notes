package server

import (
	"github.com/gin-gonic/gin"
)

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, msg)
}
