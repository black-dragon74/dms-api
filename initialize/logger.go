package initialize

import (
	"github.com/black-dragon74/dms-api/config"
	"go.uber.org/zap"
	"log"
)

func Logger(c *config.Config) *zap.Logger {
	var lgr *zap.Logger
	var err error

	if c.GetEnv() == "dev" {
		lgr, err = zap.NewDevelopmentConfig().Build()
	} else {
		lgr, err = zap.NewProductionConfig().Build()
	}

	if err != nil {
		log.Fatal("Unable to initialize the logger", err)
	}

	lgr.Info("[Initialize] [Logger] Loaded config and logger, bootstrap begin")
	return lgr
}
