package app

import (
	"context"
	"github.com/black-dragon74/dms-api/app/router"
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	serverAddress = "127.0.0.1:8000"
)

func Start(lgr *zap.Logger) {
	rtr := router.NewRouter(lgr)

	srv := &http.Server{
		Handler: handlers.RecoveryHandler()(rtr),
		Addr:    serverAddress,
	}

	go gracefulShutdown(lgr, srv)

	lgr.Sugar().Info("Server is up and running on http://", serverAddress)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		lgr.Error("Failed to start the server")
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
