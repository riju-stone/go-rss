package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/riju-stone/go-rss/internal/database"
	"github.com/riju-stone/go-rss/utils"
)

type FeedParams struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FeedModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	Name      string    `json:"feed_name"`
	UserID    uuid.UUID `json:"user_id"`
}

func FormatFeedModel(feed database.Feed) FeedModel {
	return FeedModel{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Url:       feed.Url,
		Name:      feed.FeedName,
		UserID:    feed.UserID,
	}
}

func HandleCreateFeed(w http.ResponseWriter, r *http.Request, dbq *database.Queries, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := FeedParams{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.ErrorResponse(w, 400, "Error parsing New Feed Payload")
		return
	}

	feed, err := dbq.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedName:  params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed to Create new RSS Feed Entry")
		return
	}

	utils.JsonResponse(w, 203, FormatFeedModel(feed))
}

func HandleGetAllFeeds(w http.ResponseWriter, r *http.Request, dbq *database.Queries) {
	feeds, err := dbq.GetAllFeeds(r.Context())
	if err != nil {
		utils.ErrorResponse(w, 400, "Could not fetch feeds")
		return
	}
	utils.JsonResponse(w, 200, feeds)
}
