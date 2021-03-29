package app

import (
	"context"
	"github.com/black-dragon74/dms-api/app/router"
	"github.com/black-dragon74/dms-api/config"
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(cfg *config.Config, lgr *zap.Logger) {
	rtr := router.NewRouter(lgr, cfg)

	srv := &http.Server{
		Handler: handlers.RecoveryHandler()(rtr),
		Addr:    cfg.API.GetAddress(),
	}

	go gracefulShutdown(lgr, srv)

	lgr.Sugar().Info("[App] [Start] Server is up and running on http://", cfg.API.GetAddress())
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		lgr.Sugar().Error("[App] [Start] Failed to start the server", err.Error())
	}
}

func gracefulShutdown(lgr *zap.Logger, srv *http.Server) {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, os.Interrupt, os.Kill)

	<-termChan
	lgr.Info("[App] [gracefulShutdown] Attempting graceful shutdown")

	ctx, cFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cFunc()

	go func() {
		_ = srv.Shutdown(ctx)
	}()
}
