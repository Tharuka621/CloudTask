package main

import (
	"log"
	"os"

	"cloudtask/task-service/internal/handlers"
	"cloudtask/task-service/internal/models"
	"cloudtask/task-service/internal/repositories"
	"cloudtask/task-service/internal/services"
	"cloudtask/task-service/pkg/database"
	"cloudtask/task-service/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database.ConnectPostgres()
	database.ConnectRedis()

	database.DB.AutoMigrate(&models.Project{}, &models.Task{}, &models.TaskComment{})

	repo := repositories.NewTaskRepository()
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })

	api := app.Group("/api/tasks")
	api.Use(middleware.Protected()) // JWT required

	api.Post("/", handler.CreateTask)
	api.Get("/", handler.GetTasks)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Task service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
