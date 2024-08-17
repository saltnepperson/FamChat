package handler

import (
	"net/http"

	rs "github.com/saltnepperson/FamChat/cmd/server/responses"	
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := HealthCheckResponse{Message: "Health check endpoint"}

	rs.JSON(w, http.StatusOK, res)
}
