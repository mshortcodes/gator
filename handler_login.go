package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("invalid number of args given")
	}

	user := cmd.args[0]
	dbUser, _ := s.db.GetUser(context.Background(), user)
	if dbUser.Name != user {
		log.Fatalf("that account doesn't exit")
	}

	err := s.cfg.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %v\n", user)
	return nil
}
