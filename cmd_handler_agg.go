package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd Command) error {
	fetchURL := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	feed, err := fetchFeed(ctx, fetchURL)
	if err != nil {
		return err
	}

	fmt.Println(feed.Channel.Title)
	fmt.Println(feed.Channel.Link)
	fmt.Println(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("%v\n", item)
	}

	return nil
}
