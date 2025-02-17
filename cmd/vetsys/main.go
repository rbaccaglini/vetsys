package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/vetsys/internal/routes"
	"github.com/rbaccaglini/vetsys/pkg/utils/logger"
)

func main() {
	logger.Info("Starting server...")

	router := gin.Default()

	routes.InitRouter(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
