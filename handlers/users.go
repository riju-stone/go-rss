package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/riju-stone/go-rss/internal/auth"
	"github.com/riju-stone/go-rss/internal/database"
	"github.com/riju-stone/go-rss/utils"
)

type UserParams struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type UserModel struct {
	ID        uuid.UUID `json:"id"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

// Function to format new user response as per UserModel
func FormatUserModel(user database.User) UserModel {
	return UserModel{
		ID:        user.ID,
		Fname:     user.Fname,
		Lname:     user.Lname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
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
		Fname:     params.Fname,
		Lname:     params.Lname,
	})
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed to create user!")
	}

	utils.JsonResponse(w, 201, FormatUserModel(user))
}

func HandleGetUser(w http.ResponseWriter, r *http.Request, dbq *database.Queries) {
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

	utils.JsonResponse(w, 200, FormatUserModel(user))
}
