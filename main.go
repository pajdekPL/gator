package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/pajdekpl/gator/internal/config"
	"github.com/pajdekpl/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error during reading config: %v", err)
	}
	db, err := sql.Open("postgres", config.DBURL)
	if err != nil {
		log.Fatalf("error opening db connection: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	appState := state{
		config: &config,
		db:     dbQueries,
	}

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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
		log.Fatalln(err)
	}
}
