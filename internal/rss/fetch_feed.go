package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

var (
	defaultTimeout = 10 * time.Second
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var rssFeed RSSFeed
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{
		Timeout: defaultTimeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()
	decoder := xml.NewDecoder(res.Body)

	err = decoder.Decode(&rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, rssItem := range rssFeed.Channel.Item {
		rssItem.Title = html.UnescapeString(rssItem.Title)
		rssItem.Description = html.UnescapeString(rssItem.Description)
		rssFeed.Channel.Item[i] = rssItem
	}
	return &rssFeed, nil
}

func PrintFeedItemsTitles(rssFeed *RSSFeed) {
	for _, rssItem := range rssFeed.Channel.Item {
		fmt.Println("* ", rssItem.Title)
	}
}
