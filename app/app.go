package app

import (
	"context"
	"fmt"
	"github.com/black-dragon74/dms-api/app/router"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/initialize"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(cfg *config.Config, lgr *zap.Logger) {
	// Init the data store
	store, err := initialize.DataStore(lgr, cfg)
	if err != nil {
		lgr.Error(fmt.Sprintf("[App] [Start] [DataStore] %s", err.Error()))
		return
	}

	// Initialize the redis store, it will do nothing ig `api.useRedis` is false
	rds := initialize.RedisStore(cfg, lgr)

	rtr := router.NewRouter(lgr, cfg, store, rds)
	if rtr == nil {
		// Error already logged in `NewRouter`
		return
	}

	srv := &http.Server{
		Handler: handlers.RecoveryHandler()(rtr),
		Addr:    cfg.API.GetAddress(),
	}

	go gracefulShutdown(cfg, lgr, srv, rds)

	lgr.Info(fmt.Sprintf("[App] [Start] Server is up and running on http://%s", cfg.API.GetAddress()))
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		lgr.Error(fmt.Sprintf("[App] [Start] Failed to start the server, %s", err.Error()))
	}
}

func gracefulShutdown(cfg *config.Config, lgr *zap.Logger, srv *http.Server, rds *redis.Client) {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, os.Interrupt, os.Kill)

	<-termChan
	lgr.Info("[App] [gracefulShutdown] Attempting graceful shutdown")

	ctx, cFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cFunc()

	go func() {
		_ = srv.Shutdown(ctx)

		if cfg.API.UseRedis() {
			_ = rds.Shutdown(ctx)
		}
	}()
}
