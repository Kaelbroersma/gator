package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kaelbroersma/gator/internal/database"
)

func handleAgg(s *state, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("agg takes 1 <arg>. duration. (1s, 1m, 1h)")
	}
	timeBetween, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetween)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("feed found. scraping in progress.")
	scrapeFeed(s.db, nextFeed)

	return nil
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched.", feed.Name)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't fetch feed at %s", feed.Url)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %v\n", item.Title)
	}
	log.Printf("Feed %s collected. %v posts found.", feed.Name, len(feedData.Channel.Item))
}
