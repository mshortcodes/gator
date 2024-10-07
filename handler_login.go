package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("invalid number of args given")
	}

	user := cmd.args[0]
	err := s.cfg.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %v\n", user)
	return nil
}
