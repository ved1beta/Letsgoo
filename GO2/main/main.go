package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Client represents a connected WebSocket client
type Client struct {
	conn *websocket.Conn
	hub  *Hub
	mu   sync.Mutex
}

// Hub manages all connected clients
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

// Create a new hub instance
func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			fmt.Printf("New client connected. Total clients: %d\n", len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.conn.Close()
			}
			h.mu.Unlock()
			fmt.Printf("Client disconnected. Total clients: %d\n", len(h.clients))

		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				client.mu.Lock()
				err := client.conn.WriteMessage(websocket.TextMessage, message)
				client.mu.Unlock()
				if err != nil {
					fmt.Printf("Error broadcasting to client: %v\n", err)
					client.conn.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client.html")
}

func handleWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	client := &Client{
		conn: conn,
		hub:  hub,
	}

	client.hub.register <- client

	// Handle incoming messages
	go func() {
		defer func() {
			client.hub.unregister <- client
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Read error:", err)
				break
			}
			fmt.Printf("Received message: %s\n", message)
			client.hub.broadcast <- message
		}
	}()
}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(hub, w, r)
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
