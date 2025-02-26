-- name: CreateFollowFeed :one

WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
) SELECT inserted_feed_follow.*, users.name as user_name, feeds.name as feed_name
FROM inserted_feed_follow INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;


-- name: GetFeedFollowsForUser :many
SELECT users.name as user_name, feeds.name as feed_name 
FROM users 
INNER JOIN feed_follows ON feed_follows.user_id = $1
INNER JOIN feeds ON feeds.id = feed_follows.feed_id;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
USING feeds
WHERE feeds.id = feed_follows.feed_id AND feed_follows.user_id = $1 AND feeds.URL = $2;