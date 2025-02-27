package main

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/status", handlers.StatusHandler)
	r.POST("/parse", handlers.ParseCodeHandler)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
