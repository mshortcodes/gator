package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 && len(cmd.args) != 1 {
		return fmt.Errorf("invalid number of args")
	}

	limit := 2
	if len(cmd.args) == 1 {
		if newLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = newLimit
		} else {
			return fmt.Errorf("invalid limit: %v", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %v", err)
	}

	fmt.Printf("found %d posts for %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("====================")
	}

	return nil
}
