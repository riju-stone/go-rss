package handlers

import (
	"net/http"

	"github.com/riju-stone/go-rss/utils"
)

func HandleServerError(w http.ResponseWriter, r *http.Request) {
	utils.ErrorResponse(w, 400, "Something Went Wrong!")
}
