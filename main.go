package main

import (
	"fmt"
	"os"

	"github.com/KiranSatyaRaj/gator-cli/internal/config"
)

func main() {
	args := os.Args
	if len(args) <= 2 {
		fmt.Errorf("Invalid num of args, expected atleast two")
		os.Exit(1)
	}

	st := &state{}
	cfg := config.Read()
	st.cfg = &cfg
	cmds := commands{}
	cmds.cmds = make(map[string]func(*state, command) error)
	cmd := command{args[1], args[2:]}
	cmds.register(args[1], handlerLogin)
	if err := cmds.run(st, cmd); err != nil {
		panic(err)
	}
}
