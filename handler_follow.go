package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/pajdekpl/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>", cmd.Name)
	}

	url, err := url.ParseRequestURI(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("incorrect url: %v", err)
	}

	feedId, err := s.db.GetFeedIdByUrl(context.Background(), url.String())

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("feed with URL: %s doesn't exist", url.String())
		}
		return err
	}

	creationTime := time.Now().UTC()

	feedFollow := database.CreateFollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
		UserID:    user.ID,
		FeedID:    feedId,
	}
	createdFeedFollow, err := s.db.CreateFollowFeed(context.Background(), feedFollow)

	if err != nil {
		return err
	}
	fmt.Printf("Feed: %+v\n", createdFeedFollow)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage %s", cmd.Name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <feed_url>", cmd.Name)
	}

	err := s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	})

	if err != nil {
		return err
	}
	return nil
}
