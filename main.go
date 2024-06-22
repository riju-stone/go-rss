package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/riju-stone/go-rss/routes/v1"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Could not locate PORT config in environment")
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

	// Initializing a new HTTP server using the declared router
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Executing the server to listen to a particular port
	log.Printf("Server Listening on Port: %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
