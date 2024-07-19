package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/riju-stone/go-rss/handlers"
)

func InitV1Routes() *chi.Mux {
	v1Router := chi.NewRouter()

	// Route to check API health
	v1Router.Get("/health", handlers.HandleHealthCheck)

	// Route to check server errors
	v1Router.Get("/error", handlers.HandleServerError)
	return v1Router
}
