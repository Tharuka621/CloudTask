package main

import (
	"log"
	"os"

	"cloudtask/notification-service/internal/handlers"
	"cloudtask/notification-service/internal/services"
	"cloudtask/notification-service/pkg/database"
	"cloudtask/notification-service/pkg/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database.ConnectRedis()

	hub := services.NewHub()
	go hub.Run()
	go services.ListenForNotifications(hub)

	app := fiber.New()
	app.Use(logger.New())

	// Middleware to check if it's a websocket request
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })

	ws := app.Group("/ws/notifications")
	ws.Use(middleware.WSProtected())
	ws.Get("/", handlers.ServeWS(hub))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	log.Printf("Notification service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
