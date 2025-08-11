package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var fetchedFeed RSSFeed

	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &fetchedFeed, err
	}

	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	fetchedData, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	err = xml.Unmarshal(fetchedData, &fetchedFeed)
	if err != nil {
		return &RSSFeed{}, err
	}

	fetchedFeed.Channel.Title = html.UnescapeString(fetchedFeed.Channel.Title)
	fetchedFeed.Channel.Description = html.UnescapeString(fetchedFeed.Channel.Description)

	for i := range fetchedFeed.Channel.Item {
		fetchedFeed.Channel.Item[i].Title = html.UnescapeString(fetchedFeed.Channel.Item[i].Title)
		fetchedFeed.Channel.Item[i].Description = html.UnescapeString(fetchedFeed.Channel.Item[i].Description)
	}

	return &fetchedFeed, nil
}

func addFeed(name string, url string) error {
	return nil
}
