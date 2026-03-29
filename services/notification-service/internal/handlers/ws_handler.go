package handlers

import (
	"log"

	"cloudtask/notification-service/internal/services"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ServeWS(hub *services.Hub) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		userIDFloat, ok := c.Locals("user_id").(float64)
		if !ok {
			log.Println("Invalid user_id in context")
			return
		}

		userID := uint(userIDFloat)
		client := &services.Client{
			UserID: userID,
			Conn:   c,
		}

		hub.Register(client)

		defer func() {
			hub.Unregister(client)
		}()

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}
			if mt == websocket.TextMessage {
				log.Printf("Received message from User %d: %s", userID, msg)
			}
		}
	})
}
