package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/riju-stone/go-rss/internal/database"
	"github.com/riju-stone/go-rss/utils"
)

type UserParams struct {
	Name string `json:"name"`
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request, dbq *database.Queries) {
	decoder := json.NewDecoder(r.Body)
	params := UserParams{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.ErrorResponse(w, 400, "Error parsing create user payload")
		return
	}

	user, err := dbq.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed to create user!")
	}

	utils.JsonResponse(w, 200, user)
}
