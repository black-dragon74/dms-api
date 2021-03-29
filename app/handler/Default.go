package handler

import (
	"go.uber.org/zap"
	"net/http"
)

func WelcomeHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lgr.Info("[Handler] [WelcomeHandler] Handling /")

		_, err := w.Write([]byte(`{"msg": "Welcome to MUJ DMS REST API2.0 by Nick", "help": "Invalid route. Please refer to the documentation"}`))
		if err != nil {
			lgr.Error(err.Error())
		}
	}
}
