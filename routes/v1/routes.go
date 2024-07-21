package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riju-stone/go-rss/handlers"
	"github.com/riju-stone/go-rss/internal/database"
)

func InitV1Routes(dbq *database.Queries) *chi.Mux {
	v1Router := chi.NewRouter()

	// Route to check API health
	v1Router.Get("/health", handlers.HandleHealthCheck)

	// Route to check server errors
	v1Router.Get("/error", handlers.HandleServerError)

	// Route to create an user
	v1Router.Post("/create-user", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleCreateUser(w, r, dbq)
	})

	return v1Router
}
