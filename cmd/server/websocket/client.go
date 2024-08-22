package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	room   string
	userID int
	chatID int
}

// ReadPump pumps messages from the websocket connection to the Hub
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error during read pump: %v", err)
			}
		}

		c.hub.broadcast <- message
	}
}

// WritePump pumps messgaes from the hub to the websocket connection
func (c *Client) WritePump() {
	ticker := time.NewTicker(54 * time.Second)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("Error while getting next writer:", err)
				return
			}

			w.Write(message)

			// Add queued chat messages to the current WebSocket message.
			n := len(c.send)

			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Println("Error while closing writer:", err)
				return
			}
		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
