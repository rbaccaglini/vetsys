package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	holiday_handler "github.com/rbaccaglini/vetsys/internal/handlers/holiday"
	holiday_repo "github.com/rbaccaglini/vetsys/internal/repositories/holiday"
	"github.com/rbaccaglini/vetsys/internal/routes"
	holiday_service "github.com/rbaccaglini/vetsys/internal/services/holiday"
	"github.com/rbaccaglini/vetsys/pkg/database/mongodb"
	"github.com/rbaccaglini/vetsys/pkg/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("Starting server...")

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect on DB", err)
		return
	}
	logger.Info("DB Connecteded")

	holidayHandler := initDependencies(database)
	router := gin.Default()

	routes.InitRouter(&router.RouterGroup, holidayHandler)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func initDependencies(database *mongo.Database) holiday_handler.HolidayHandlerInterface {
	repo := holiday_repo.NewHolidayRepository(database)
	service := holiday_service.NewHolidayDomainService(repo)
	return holiday_handler.NewUserHandlerInterface(service)
}
