package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kaelbroersma/gator/internal/config"
	"github.com/kaelbroersma/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Couldn't read saved Config file. \nError: %v\n", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBurl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	userState := &state{cfg: &cfg, db: dbQueries}

	cmdRegistry := Commands{
		handlers: make(map[string]handlerFunc),
	}

	cmdRegistry.register("login", HandleLogin)
	cmdRegistry.register("register", handleRegister)
	cmdRegistry.register("reset", resetHandler)
	cmdRegistry.register("users", handleListUsers)
	cmdRegistry.register("agg", handleAgg)
	cmdRegistry.register("addfeed", handleAddFeed)
	cmdRegistry.register("feeds", handleListFeeds)

	if len(os.Args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmdRegistry.run(userState, Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
