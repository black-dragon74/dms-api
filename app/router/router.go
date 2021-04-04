package router

import (
	"github.com/black-dragon74/dms-api/app/handler"
	"github.com/black-dragon74/dms-api/app/middleware"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/types"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter(lgr *zap.Logger, cfg *config.Config, store *types.DataStoreModel) *mux.Router {
	rtr := mux.NewRouter()
	rtr.Use(middleware.WithContentJSON)

	// Map route to handler functions
	// Default Handler
	rtr.HandleFunc("/", handler.WelcomeHandler(lgr)).Methods(http.MethodGet)

	// Routes without session ID, require respective data stores
	rtr.HandleFunc("/mess_menu", handler.MessMenuHandler(lgr, store.MessMenuData)).Methods(http.MethodGet)
	rtr.HandleFunc("/contacts", handler.ContactsHandler(lgr, store.ContactsData)).Methods(http.MethodGet)

	// Routes part of auth handshake
	rtr.HandleFunc("/captcha", handler.GetCaptchaHandler(lgr)).Methods(http.MethodGet)
	rtr.HandleFunc(
		"/captcha_auth",
		handler.CaptchaAuthHandler(lgr)).Methods(http.MethodGet)

	// Routes that need a session ID to through
	rtr.HandleFunc("/dashboard", handler.DashboardHandler(lgr)).Methods(http.MethodGet)
	rtr.HandleFunc("/attendance", handler.AttendanceHandler(lgr)).Methods(http.MethodGet)
	rtr.HandleFunc("/results", handler.ResultsHandler(lgr)).Methods(http.MethodGet)
	rtr.HandleFunc("/internals", handler.InternalsHandler(lgr)).Methods(http.MethodGet)
	rtr.HandleFunc("/gpa", handler.GPAHandler(lgr)).Methods(http.MethodGet)

	// Some browsers request for favicon even when the content type is set to JSON, handle that
	rtr.HandleFunc("/favicon.ico", handler.FaviconHandler(lgr)).Methods(http.MethodGet)

	if cfg.GetEnv() == config.DevEnv {
		rtr.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			// Write your route specific tests here

		}).Methods(http.MethodGet)
	}

	// Should be the last to allow sieving of requests
	rtr.HandleFunc("/{*}", handler.FourOhFourHandler(lgr)).Methods(http.MethodGet)

	return rtr
}
