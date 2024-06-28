package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/riju-stone/go-rss/routes/v1"
	"github.com/riju-stone/go-rss/utils"
)

func main() {
	// Load env variables
	godotenv.Load(".env")

	// Creating customer logger directory and file
	logFile := os.Getenv("LOGFILE")
	if dirError := os.Mkdir("logs", os.ModePerm); dirError != nil {
		panic(dirError)
	}

	f, fileError := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if fileError != nil {
		panic(fileError)
	}

	// Initializing Custom Logger
	rssLog := utils.InitLogger(f)
	defer f.Close()

	port := os.Getenv("PORT")
	if port == "" {
		rssLog.Error("Could not locate PORT config in environment")
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
	rssLog.Debug("V1 Routes Mounted")

	// Initializing a new HTTP server using the declared router
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Executing the server to listen to a particular port
	rssLog.Info("Server Listening on Port: ", port)
	serverError := server.ListenAndServe()
	if serverError != nil {
		rssLog.Panic(serverError)
	}
}
