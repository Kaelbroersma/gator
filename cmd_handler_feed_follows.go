package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaelbroersma/gator/internal/database"
)

func handleFollowing(s *state, cmd Command, user database.User) error {
	ctx := context.Background()

	followedFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println("======================================")
	fmt.Printf("%v's followed feeds", s.cfg.CurrentUser)
	fmt.Println()
	for _, feed := range followedFeeds {
		fmt.Println(feed)
	}
	fmt.Println("======================================")

	return nil
}

func handleFollow(s *state, cmd Command, user database.User) error {
	ctx := context.Background()

	if len(cmd.Args) < 1 {
		return fmt.Errorf("Must call follow with atleast 1 <args>")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	newRow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Successfully followed feed %v\nFor user %v", newRow.FeedName, newRow.UserName)
	fmt.Println()

	return nil
}

func handleUnfollow(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please provide a URL to unfollow")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("\nSuccessfully unfollowed %v\n", url)

	return nil
}
