package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kaelbroersma/gator/internal/database"
)

func handleBrowse(s *state, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if declaredLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = declaredLimit
		} else {
			return fmt.Errorf("Invalid limit: %v", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Println("=====================================================")
	fmt.Printf("Found %v posts for user %v\n", len(posts), user.Name)
	fmt.Println("=====================================================")
	for _, post := range posts {
		fmt.Printf("\n[+]%s\n", post.Title)
		fmt.Printf("%s\n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Println()
		fmt.Printf("%v\n", post.Description)
		fmt.Println()
		fmt.Printf("Find it at %v\n", post.Url)
		fmt.Println()
		fmt.Println("====================================================")
	}

	return nil
}
