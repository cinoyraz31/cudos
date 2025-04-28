package main

import (
	"cudo_task_service/config"
	"cudo_task_service/exceptions"
	"cudo_task_service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	Init()
	db := config.OpenConnection()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(exceptions.ErrorHandlerInternalServerError)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to cudo interview")
	})

	routes.TranscationRoutes(app, db)

	err := app.Listen(os.Getenv("APP_URL"))
	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Failed to close database connection!", err)
		}
		sqlDB.Close()
	}()
}
