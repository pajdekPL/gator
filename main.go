package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pajdekpl/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error during reading config %v", err)
	}
	appState := state{
		config: &config,
	}
	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("please give at least single command")
	}
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(&appState, cmd)
	if err != nil {
		log.Fatalf("error executing cmd: %v", err)
	}
	fmt.Printf("config %+v\n", *appState.config)
}
