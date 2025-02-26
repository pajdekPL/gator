// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feeds_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFollowFeed = `-- name: CreateFollowFeed :one

WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, created_at, updated_at, user_id, feed_id
) SELECT inserted_feed_follow.id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at, inserted_feed_follow.user_id, inserted_feed_follow.feed_id, users.name as user_name, feeds.name as feed_name
FROM inserted_feed_follow INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id
`

type CreateFollowFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFollowFeedRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	UserName  string
	FeedName  string
}

func (q *Queries) CreateFollowFeed(ctx context.Context, arg CreateFollowFeedParams) (CreateFollowFeedRow, error) {
	row := q.db.QueryRowContext(ctx, createFollowFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFollowFeedRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.UserName,
		&i.FeedName,
	)
	return i, err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT users.name as user_name, feeds.name as feed_name 
FROM users 
INNER JOIN feed_follows ON feed_follows.user_id = $1
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
`

type GetFeedFollowsForUserRow struct {
	UserName string
	FeedName string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(&i.UserName, &i.FeedName); err != nil {
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

const unfollowFeed = `-- name: UnfollowFeed :exec
DELETE FROM feed_follows
USING feeds
WHERE feeds.id = feed_follows.feed_id AND feed_follows.user_id = $1 AND feeds.URL = $2
`

type UnfollowFeedParams struct {
	UserID uuid.UUID
	Url    string
}

func (q *Queries) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, arg.UserID, arg.Url)
	return err
}
