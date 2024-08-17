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

