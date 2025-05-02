package main

import (
	"backend/Handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"time"
)

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/status", Handlers.StatusHandler)
	r.POST("/parse", Handlers.ParseCodeHandler)
	r.POST("/discs", Handlers.DiscsHandler)

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
