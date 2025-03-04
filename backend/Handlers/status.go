package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
