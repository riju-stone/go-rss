package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	log "github.com/riju-stone/go-rss/logging"
	"github.com/riju-stone/go-rss/routes/v1"
)

func main() {
	// Load env variables
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Error("Could not locate PORT config in environment")
	}

	// Creating a new router
	router := chi.NewRouter()

	// Adding cors config to the global router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Adding & Mounting Sub-routes
	v1Router := routes.InitV1Routes()
	router.Mount("/v1", v1Router)
	log.Debug("V1 Routes Mounted")

	// Initializing a new HTTP server using the declared router
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Executing the server to listen to a particular port
	log.Info("Server Listening on Port: %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("Failed to initialize server: %v", err)
	}
}
