package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("no name provided")
	}

	name := cmd.args[0]
	user, _ := s.db.GetUser(context.Background(), name)
	if user.Name == name {
		log.Fatal("user already exists")
	}

	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %v", err)
	}

	err = s.cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %v", err)
	}

	printUser(dbUser)
	return nil
}

func printUser(user database.User) {
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("ID: %s\n", user.ID)
}
