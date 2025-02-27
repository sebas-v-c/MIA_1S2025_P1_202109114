package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeRequest struct {
	Code string `json:"code"`
}

func ParseCodeHandler(c *gin.Context) {
	// TODO: parse code and return the result
	var req CodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	fmt.Println("Received req: ", req.Code)
	c.JSON(http.StatusOK, gin.H{"message": "Code received", "code": req.Code})
}
