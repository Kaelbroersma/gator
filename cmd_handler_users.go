package main

import (
	"context"
	"fmt"
	"time"
	"log"

	"github.com/google/uuid"
	"github.com/kaelbroersma/gator/internal/database"
)

func handleRegister(s *state, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Must input name with command register")
	}

	name := cmd.Args[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		return fmt.Errorf("User already exists in database.")
	}

	userStruct := database.User{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}

	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams(userStruct))
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully registered user %s with UUID %v\n", userStruct.Name, newUser.ID)

	return nil
}


func HandleLogin(s *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username.")
	}

	name := cmd.Args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		log.Fatal("user does not exist.")
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Unable to set user in login handler. %v", err)
	}

	fmt.Println("User has been set to " + name)

	return nil
}

func handleListUsers (s *state, cmd Command) error {
	ctx := context.Background()
	currentUser := s.cfg.CurrentUser

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %v (current)\n", currentUser)
			continue
		} 
		fmt.Printf("* %v\n",user.Name)
	}
	return nil
}



