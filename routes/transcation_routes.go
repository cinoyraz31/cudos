package routes

import (
	"cudo_task_service/controllers"
	"cudo_task_service/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TranscationRoutes(app *fiber.App, db *gorm.DB) {
	transactionRepository := repositories.NewTransactionRepository()
	transactionController := controllers.NewTransactionController(db, transactionRepository)

	app.Get("/api/v1/fraud-detection", transactionController.FraudDetection)
}
