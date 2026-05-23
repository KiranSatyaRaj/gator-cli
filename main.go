package main

import (
	"fmt"

	"github.com/KiranSatyaRaj/gator-cli/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser()
	fmt.Println(config.Read())
}
