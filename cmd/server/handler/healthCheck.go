package handler

import (
	"net/http"

	rs "github.com/saltnepperson/FamChat/pkg/responses"	
)

type HealthCheck struct {
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := HealthCheck{Message: "Health check endpoint"}

	rs.RespondWithJSON(w, http.StatusOK, res)
}
