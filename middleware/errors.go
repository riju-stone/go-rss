package middleware

import (
	"net/http"

	log "github.com/riju-stone/go-rss/logging"
	"github.com/riju-stone/go-rss/utils"
)

func HandleServerError(w http.ResponseWriter, r *http.Request) {
	log.Debug("[Request Received] Method=", r.Method, " | URI=", r.RequestURI)
	utils.ErrorResponse(w, 400, "Something Went Wrong!")
}
