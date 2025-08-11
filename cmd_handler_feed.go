package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaelbroersma/gator/internal/database"
)

func handleAddFeed(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("<args1> must be URL followed by <args2> name")
	}
	if len(cmd.Args) < 2 {
		return fmt.Errorf("<args2> must be name")
	}

	feedURL := cmd.Args[1]
	feedName := cmd.Args[0]
	ctx := context.Background()

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      feedName,
		Url:       feedURL,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error creating feed following on addfeed: %v", err)
	}

	fmt.Println()
	fmt.Println("Feed created successfully!")
	fmt.Println("=========================================================")
	fmt.Println()
	printFeed(feed)
	fmt.Println()
	fmt.Println("=========================================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID %v\n", feed.ID)
	fmt.Printf("* Created: %v\n", feed.CreatedAt)
	fmt.Printf("* Updated: %v\n", feed.UpdatedAt)
	fmt.Printf("* Name: %v\n", feed.Name)
	fmt.Printf("* URL: %v\n", feed.Url)
	fmt.Printf("* User ID: %v\n", feed.UserID)
}

func handleListFeeds(s *state, cmd Command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("---BEGIN RECORDS---")
	fmt.Println()

	for i := range feeds {
		fmt.Printf("Record number: %d\n", i)
		fmt.Printf("Name: %v\n", feeds[i].Name)
		fmt.Printf("URL: %v\n", feeds[i].Url)
		fmt.Printf("Feed owner: %v\n", feeds[i].UserName)
		fmt.Println()
	}

	fmt.Println("---END RECORDS---")

	return nil
}
