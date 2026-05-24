package main

import (
	"github.com/KiranSatyaRaj/gator-cli/internal/config"
	"github.com/KiranSatyaRaj/gator-cli/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
