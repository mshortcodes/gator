package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerGetFeedFollows(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %v", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("no feed follows found for this user")
		return nil
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}

	return nil
}
