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
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type UserModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	ApiKey    string    `json:"api_key"`
	ID        uuid.UUID `json:"id"`
}

type PostModel struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
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

func FormatPostModel(post database.Post) PostModel {
	return PostModel{
		ID:          post.ID,
		Title:       post.Title,
		Description: &post.Description.String,
		Url:         post.Url,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
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

func HandleGetUser(w http.ResponseWriter, r *http.Request, dbq *database.Queries, user database.User) {
	utils.JsonResponse(w, 200, FormatUserModel(user))
}

func ConvertToPostList(posts []database.Post) []PostModel {
	postsArr := []PostModel{}
	for _, post := range posts {
		postsArr = append(postsArr, FormatPostModel(post))
	}

	return postsArr
}

func HandleGetPostsFollowedByUser(w http.ResponseWriter, r *http.Request, dbq *database.Queries, user database.User) {
	posts, err := dbq.GetUserFollowedPosts(r.Context(), database.GetUserFollowedPostsParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed to fetch latest posts")
		return
	}

	utils.JsonResponse(w, 200, ConvertToPostList(posts))
}
