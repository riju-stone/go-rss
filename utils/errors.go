package utils

import (
	"net/http"

	log "github.com/riju-stone/go-rss/logging"
)

type errResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, mssg string) {
	// Errors below 499 are generally client side errors
	if statusCode > 499 {
		log.Error("Responding with 5XX server error")
	}

	JsonResponse(w, statusCode, errResponse{Error: mssg})
}
