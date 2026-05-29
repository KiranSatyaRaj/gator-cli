package main

import (
	"context"
	"errors"
	"fmt"
	"html"
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

func handlerAgg(state *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(html.UnescapeString(rssFeed.Channel.Title))
	fmt.Println(html.UnescapeString(rssFeed.Channel.Description))
	rssItem := rssFeed.Channel.Item
	for _, item := range rssItem {
		fmt.Println(html.UnescapeString(item.Title))
		fmt.Println(html.UnescapeString(item.Description))
	}
	return nil
}

func handlerAddFeed(state *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("not enough args")
	}
	currentUser, err := state.db.GetUser(context.Background(), state.cfg.CurrentUserName)
	feedArgs := database.CreateFeedParams{
		Name:   cmd.args[0],
		Url:    cmd.args[1],
		UserID: currentUser.ID,
	}

	feed, err := state.db.CreateFeed(context.Background(), feedArgs)
	if err != nil {
		return err
	}
	fmt.Printf("Name of the feed: %s\n", feed.Name)
	fmt.Printf("Url of the feed: %s\n", feed.Url)
	return nil
}

func handlerListFeeds(state *state, cmd command) error {
	feeds, err := state.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		username, err := state.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Name of the feed: %s, Url of the feed: %s, user: %s\n", feed.Name, feed.Url, username)
	}
	return nil
}
