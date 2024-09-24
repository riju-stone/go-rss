package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/riju-stone/go-rss/logging"
)

func JsonResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Error("Failed to parse server response: ")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)

	log.Info("Response Payload: %s", data)
}
