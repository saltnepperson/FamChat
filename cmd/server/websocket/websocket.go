package websocket

import (
    "net/http"
    "log"
    "github.com/gorilla/websocket"
)

// Upgrader configures the WebSocket upgrade options.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        // allow all origins, for now 
        return true
    },
}

// ServeWS handles WebSocket requests from clients.
func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading to WebSocket:", err)
        return
    }

    client := &Client{
        hub:  hub,
        conn: conn,
        send: make(chan []byte, 256),
    }

    hub.register <- client

    // Start the client's read and write pumps
    go client.WritePump()
    go client.ReadPump()
}

