package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KiranSatyaRaj/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(state *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("expected username, got zero args")
	}
	user, err := state.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		os.Exit(1)
	}
	state.cfg.SetUser(user.Name)
	fmt.Print("user id created\n")
	return nil
}

func handlerRegister(state *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("expected username, got zero args")
	}
	userInfo := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	user, err := state.db.CreateUser(context.Background(), userInfo)
	if err != nil {
		os.Exit(1)
	}
	state.cfg.SetUser(user.Name)
	fmt.Printf("Registered %v as user at %v\n", user.Name, user.CreatedAt)
	return nil
}

func handlerReset(state *state, cmd command) error {
	if err := state.db.DeleteAllUsers(context.Background()); err != nil {
		log.Fatal("Unable to reset users")
		return err
	}
	log.Println("Deleted All users")
	return nil
}

func handlerUsers(state *state, cmd command) error {
	users, err := state.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal("No users found")
	}
	for _, user := range users {
		if user == state.cfg.CurrentUserName {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}
	return nil
}
