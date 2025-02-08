-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, feed_name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: FetchLatestFeeds :many
SELECT * FROM feeds
ORDER BY fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFetchedFeed :one
UPDATE feeds
SET fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
