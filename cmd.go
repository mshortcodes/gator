package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	regCmds map[string]func(*state, command) error
}

// registers a cmd in the commands.regCmds map
func (c *commands) register(name string, f func(*state, command) error) {
	c.regCmds[name] = f
}

// executes a registered cmd from the commands.regCmds map
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.regCmds[cmd.name]
	if !ok {
		return fmt.Errorf("command doesn't exist")
	}

	err := f(s, cmd)
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}

	return nil
}
