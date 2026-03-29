package services

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Hub struct {
	clients    map[uint]*websocket.Conn
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mutex      sync.Mutex
}

type Client struct {
	UserID uint
	Conn   *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uint]*websocket.Conn),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client.UserID] = client.Conn
			h.mutex.Unlock()
			log.Printf("Client connected: User %d", client.UserID)
		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client.UserID]; ok {
				delete(h.clients, client.UserID)
				client.Conn.Close()
			}
			h.mutex.Unlock()
			log.Printf("Client disconnected: User %d", client.UserID)
		case message := <-h.broadcast:
			h.mutex.Lock()
			for _, conn := range h.clients {
				conn.WriteMessage(websocket.TextMessage, message)
			}
			h.mutex.Unlock()
		}
	}
}

func (h *Hub) BroadcastToUser(userID uint, message []byte) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if conn, ok := h.clients[userID]; ok {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}
