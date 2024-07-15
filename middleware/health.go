package middleware

import (
	"net/http"

	"github.com/riju-stone/go-rss/utils"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, 200, struct{}{})
}
