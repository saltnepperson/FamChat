package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

func RouteService() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	return r
}
