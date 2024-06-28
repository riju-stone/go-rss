package middleware

import (
	"net/http"

	log "github.com/riju-stone/go-rss/logging"
	"github.com/riju-stone/go-rss/utils"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Debug("[Request Received] Method=", r.Method, " | URI=", r.RequestURI)
	utils.JsonResponse(w, 200, struct{}{})
}
