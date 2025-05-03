package Handlers

import (
	"backend/Classes/Env"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	if Env.CurrentUser == nil {
		c.JSON(401, gin.H{"error": "No user is currently logged in"})
		return
	}

	c.JSON(200, Env.CurrentUser)
}
