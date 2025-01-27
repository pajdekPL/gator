package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/pajdekpl/gator/internal/database"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage %s", cmd.Name)
	}
	feeds, err := s.db.GetFeedsWithUserName(context.Background())

	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("* Feed name: 			%s\n", feed.Name)
		fmt.Printf("* Feed URL: 			%s\n", feed.Url)
		fmt.Printf("* Added by: 			%s\n", feed.CreatedBy)

	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <feed_name> <url>", cmd.Name)
	}

	url, err := url.ParseRequestURI(cmd.Args[1])

	if err != nil {
		return fmt.Errorf("incorrect url: %v", err)
	}
	feedName := cmd.Args[0]
	userUUID, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("problem with getting current user uuid %v", err)
	}

	creationTime := time.Now().UTC()
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
		Name:      feedName,
		Url:       url.String(),
		UserID:    userUUID.ID,
	}

	createdFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return err
	}

	feedFollow := database.CreateFollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
		UserID:    userUUID.ID,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFollowFeed(context.Background(), feedFollow)

	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully:")
	printFeed(createdFeed)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
