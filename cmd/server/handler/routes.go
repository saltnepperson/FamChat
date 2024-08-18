package handler

import (
	"net/http"
	"github.com/saltnepperson/FamChat/cmd/server/websocket"

	"github.com/gorilla/mux"
)

func RouteService() http.Handler {
	mux := mux.NewRouter()
	hub := websocket.NewHub()

	go hub.Run()

	// Register routes
	mux.HandleFunc("/health_check", HealthCheck)
	
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        websocket.ServeWS(hub, w, r)
  })
	
	return mux
}

