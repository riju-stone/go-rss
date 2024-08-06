package middleware

import (
	"net/http"

	"github.com/riju-stone/go-rss/internal/auth"
	"github.com/riju-stone/go-rss/internal/database"
	"github.com/riju-stone/go-rss/utils"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func AuthMiddlware(handler AuthHandler, dbq *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.AuthenticateApiKey(r.Header)
		if err != nil {
			utils.ErrorResponse(w, 401, "Failed to authenticate user")
			return
		}

		user, err := dbq.GetUserFromApiKey(r.Context(), apiKey)
		if err != nil {
			utils.ErrorResponse(w, 404, "User not found")
			return
		}

		handler(w, r, user)
	}
}
