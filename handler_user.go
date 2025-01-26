package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pajdekpl/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	username := cmd.Args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	}

	userDb, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("couldn't create user %v", err)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set user %v", err)
	}

	fmt.Println("User has been registered:")
	printUser(userDb)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)

	if err == sql.ErrNoRows {
		return fmt.Errorf("user: %s doesn't exist", username)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User has been set:", username)
	return nil
}

func printUser(user database.User) {
	fmt.Printf("User name: 		%v\n", user.Name)
	fmt.Printf("User UUID: 		%v\n", user.ID)

}
