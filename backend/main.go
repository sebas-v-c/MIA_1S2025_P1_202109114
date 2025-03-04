package main

import (
	"backend/Handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/status", Handlers.StatusHandler)
	r.POST("/parse", Handlers.ParseCodeHandler)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
