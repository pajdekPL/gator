package main

import (
	"context"
	"fmt"

	"github.com/pajdekpl/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
