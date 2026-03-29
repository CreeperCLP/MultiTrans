package main

import (
	"github.com/gin-gonic/gin"
	"clipboard-transfer-station/backend/internal/api"
)

func main() {
	r := gin.Default()
	api.SetupRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
