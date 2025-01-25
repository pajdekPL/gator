package main

import (
	"fmt"

	"github.com/pajdekpl/gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the username positional argument is required for login, example of usage: login <username>")
	}
	login := cmd.Args[0]
	err := s.config.SetUser(login)
	if err != nil {
		return err
	}
	fmt.Println("User has been set:", login)
	return nil
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
