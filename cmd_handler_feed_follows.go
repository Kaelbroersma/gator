package main

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/kaelbroersma/gator/internal/database"
)

func handleFollowing(s *state, cmd Command) error {
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUser)
	if err != nil {
		return err
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)

	fmt.Println("======================================")
	fmt.Printf("%v's followed feeds", s.cfg.CurrentUser)
	fmt.Println()
	for _, feed := range followedFeeds {
		fmt.Println(feed)
	}
	fmt.Println("======================================")
	
	return nil
}

func handleFollow(s *state, cmd Command) error {
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUser)
	if err != nil {
		return err
	}

	if len(cmd.Args) < 1 {
		return fmt.Errorf("Must call follow with atleast 1 <args>")
	}

	url := cmd.Args[0]
	feedID, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	newRow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feedID,
	})
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Successfully followed feed %v\nFor user %v", newRow.FeedName, newRow.UserName)
	fmt.Println()

	return nil
}


