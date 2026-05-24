package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/KiranSatyaRaj/gator-cli/internal/config"
	"github.com/KiranSatyaRaj/gator-cli/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Errorf("Invalid num of args, expected atleast two")
		os.Exit(1)
	}

	st := &state{}
	cfg := config.Read()
	st.cfg = &cfg
	db, err := sql.Open("postgres", st.cfg.DB_URL)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)
	st.db = dbQueries
	cmds := commands{}
	cmds.cmds = make(map[string]func(*state, command) error)
	cmd := command{args[1], args[2:]}
	switch args[1] {
	case "login":
		cmds.register("login", handlerLogin)
	case "register":
		cmds.register("register", handlerRegister)
	case "reset":
		cmds.register("reset", handlerReset)
	case "users":
		cmds.register("users", handlerUsers)
	}
	if err := cmds.run(st, cmd); err != nil {
		panic(err)
	}
}
