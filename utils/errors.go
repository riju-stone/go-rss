package utils

import (
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, mssg string) {
	// Errors below 499 are generally client side errors
	if statusCode > 499 {
		log.Println("Responding with 5XX server error: ", mssg)
	}

	JsonResponse(w, statusCode, errResponse{Error: mssg})
}
