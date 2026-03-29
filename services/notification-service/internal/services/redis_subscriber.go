package services

import (
	"encoding/json"
	"log"

	"cloudtask/notification-service/pkg/database"
)

type Notification struct {
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}

func ListenForNotifications(hub *Hub) {
	pubsub := database.RedisClient.Subscribe(database.Ctx, "notifications")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		var notif Notification
		if err := json.Unmarshal([]byte(msg.Payload), &notif); err != nil {
			log.Printf("failed to unmarshal notification payload: %v", err)
			continue
		}

		log.Printf("Broadcasting to user %d: %s", notif.UserID, notif.Message)
		hub.BroadcastToUser(notif.UserID, []byte(notif.Message))
	}
}
