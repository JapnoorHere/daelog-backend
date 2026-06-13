package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": data})
}

func BadRequest(c *gin.Context, code, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   gin.H{"code": code, "message": message},
	})
}

func InternalError(c *gin.Context, code, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   gin.H{"code": code, "message": message},
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   gin.H{"code": "NOT_FOUND", "message": message},
	})
}
