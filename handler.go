package main

import (
	"errors"
	"fmt"
)

func handlerLogin(state *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("expected username, got zero args")
	}
	state.cfg.SetUser(cmd.args[0])
	fmt.Print("user id created\n")
	return nil
}
