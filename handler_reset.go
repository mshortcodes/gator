package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("couldn't delete users: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully deleted all users")
	return nil
}
