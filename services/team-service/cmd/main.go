package main

import (
	"log"
	"os"

	"cloudtask/team-service/internal/handlers"
	"cloudtask/team-service/internal/models"
	"cloudtask/team-service/internal/repositories"
	"cloudtask/team-service/internal/services"
	"cloudtask/team-service/pkg/database"
	"cloudtask/team-service/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // ignore error, default to environment vars

	// Database initialization
	database.ConnectPostgres()

	// Auto Migration
	database.DB.AutoMigrate(&models.Team{}, &models.TeamMember{})

	// Dependency Injection
	repo := repositories.NewTeamRepository()
	service := services.NewTeamService(repo)
	handler := handlers.NewTeamHandler(service)

	app := fiber.New()
	app.Use(logger.New())

	// Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api/teams")
	api.Use(middleware.Protected()) // JWT middleware for all team routes

	api.Post("/", handler.CreateTeam)
	api.Get("/", handler.GetMyTeams)
	api.Post("/:id/members", handler.AddMember)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	log.Printf("Team service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
