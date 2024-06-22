package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/riju-stone/go-rss/middleware"
)

func InitV1Routes() *chi.Mux {
	v1Router := chi.NewRouter()

	// Route to check API health
	v1Router.Get("/health", middleware.HandleHealthCheck)

	// Route to check server errors
	v1Router.Get("/error", middleware.HandleServerError)
	return v1Router
}
