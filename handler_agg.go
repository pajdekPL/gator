package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pajdekpl/gator/internal/database"
	"github.com/pajdekpl/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <time_between_reqs>\n time_between_reqs - for example: 1s, 1m, 1h", cmd.Name)
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapFeeds(s)
	}
}

func scrapFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("problem with getting next feed %v\n", err)
		return
	}
	fmt.Printf("fetching feed: %v\n", feed.Name)
	scrapFeed(s.db, feed)
}

func scrapFeed(db *database.Queries, feed database.Feed) {
	feedContent, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("problem with fetching feed content %v\n", err)
		return
	}

	updateTime := time.Now().UTC()
	err = db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: updateTime,
	})

	if err != nil {
		fmt.Printf("problem with marking feed as fetched %v\n", err)
		return
	}

	for _, rssItem := range feedContent.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, rssItem.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: updateTime,
			UpdatedAt: updateTime,
			Title:     rssItem.Title,
			Url:       rssItem.Link,
			Description: sql.NullString{
				String: rssItem.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
}
