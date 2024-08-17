package handler

import (
	"net/http"
)

func RouteService() *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health_check", HealthCheck)
	
	return mux
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy server"))
}
