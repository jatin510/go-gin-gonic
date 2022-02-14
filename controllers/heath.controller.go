package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h *HealthController) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Working !",
	})
}
