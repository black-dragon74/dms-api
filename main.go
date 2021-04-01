package main

import (
	"errors"
	"github.com/black-dragon74/dms-api/app"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/initialize"
	"os"
)

func main() {
	// Check for required files
	// No error handling as the function panics when file is AWOL
	checkForDeps()

	// Load the app config
	cfg := config.Load()

	// Initialize the logger
	lgr := initialize.Logger(cfg)

	// Mount the router and start the server
	app.Start(cfg, lgr)
}

func checkForDeps() {
	deps := []string{
		"config.toml",
		"data/mess_menu.json",
		"data/faculties.json",
	}

	for _, v := range deps {
		if _, err := os.Stat(v); errors.Is(err, os.ErrNotExist) {
			panic("A required dependency " + v + " is missing.")
		}
	}
}
