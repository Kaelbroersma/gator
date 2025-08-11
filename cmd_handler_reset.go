package main

import(
	"context"
	"fmt"
	"log"
)

func resetHandler(s *state, cmd Command) error {
	ctx := context.Background()
	err := s.db.ResetUsers(ctx)
	if err != nil{
		log.Fatal("Database reset failed.\n")
	}
	fmt.Println("Database reset successful.")
	return nil
}

