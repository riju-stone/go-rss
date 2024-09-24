// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: followers.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollower = `-- name: CreateFeedFollower :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowerParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFeedFollower(ctx context.Context, arg CreateFeedFollowerParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollower,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const getFollowedFeeds = `-- name: GetFollowedFeeds :many
SELECT id, created_at, updated_at, user_id, feed_id FROM feed_follows WHERE user_id=$1
`

func (q *Queries) GetFollowedFeeds(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowedFeeds, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}