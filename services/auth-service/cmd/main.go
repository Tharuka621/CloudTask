package main

import (
	"log"
	"os"

	"cloudtask/auth-service/internal/handlers"
	"cloudtask/auth-service/internal/models"
	"cloudtask/auth-service/internal/repositories"
	"cloudtask/auth-service/internal/services"
	"cloudtask/auth-service/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // ignore error, default to environment vars

	// Database initialization
	database.ConnectPostgres()
	database.ConnectRedis()

	// Auto Migration
	database.DB.AutoMigrate(&models.User{}, &models.RefreshToken{})

	// Dependency Injection
	repo := repositories.NewUserRepository()
	service := services.NewAuthService(repo)
	handler := handlers.NewAuthHandler(service)

	app := fiber.New()
	app.Use(logger.New())

	// Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api/auth")
	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Auth service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
