package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(s *state, cmd command) error) {
	c.registeredCommands[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("given cmd: %s doesn't exist", cmd.Name)
	}
	return handler(s, cmd)
}
