
-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, URL, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedsWithUserName :many
SELECT feeds.name, feeds.URL, users.name as created_by FROM
feeds INNER JOIN 
users ON feeds.user_id = users.id;

-- name: GetFeedIdByUrl :one
SELECT id FROM feeds WHERE URL = $1;
