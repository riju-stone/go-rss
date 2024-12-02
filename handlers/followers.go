package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/riju-stone/go-rss/internal/database"
	"github.com/riju-stone/go-rss/utils"
)

type FollowFeedParams struct {
	FeedID uuid.UUID `json:"feed_id"`
}

type GetFollowedFeedsParams struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type FeedFollowerModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func FormatFeedFollowerModel(follower database.FeedFollow) FeedFollowerModel {
	return FeedFollowerModel{
		ID:        follower.ID,
		CreatedAt: follower.CreatedAt,
		UpdatedAt: follower.UpdatedAt,
		UserID:    follower.UserID,
		FeedID:    follower.FeedID,
	}
}

func HandleFollowFeed(w http.ResponseWriter, r *http.Request, dbq *database.Queries, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := FollowFeedParams{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed parse request payload: %s", err.Error())
		return
	}

	follower, err := dbq.CreateFeedFollower(r.Context(), database.CreateFeedFollowerParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed to follow feed : %s", err.Error())
		return
	}

	utils.JsonResponse(w, 201, FormatFeedFollowerModel(follower))
}

func HandleGetFollowedFeeds(w http.ResponseWriter, r *http.Request, dbq *database.Queries, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := GetFollowedFeedsParams{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.ErrorResponse(w, 400, "Failed to parse request payload: %s", err.Error())
	}

	feeds, err := dbq.GetFollowedFeeds(r.Context(), params.UserID)
	if err != nil {
		utils.ErrorResponse(w, 400, "Error occured while trying to fetch followed feeds: %s", err.Error())
	}

	utils.JsonResponse(w, 200, feeds)
}
