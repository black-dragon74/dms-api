package main

import (
	"github.com/black-dragon74/dms-api/app"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/initialize"
)

func main() {
	cfg := config.Load()
	lgr := initialize.Logger(cfg)
	app.Start(cfg, lgr)
}
