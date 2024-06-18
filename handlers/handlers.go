package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck to check if the server is running
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
