package router

import (
	"github.com/black-dragon74/dms-api/app/handler"
	"github.com/black-dragon74/dms-api/app/middleware"
	"github.com/black-dragon74/dms-api/config"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter(lgr *zap.Logger, cfg *config.Config) *mux.Router {
	rtr := mux.NewRouter()
	rtr.Use(middleware.WithContentJSON)

	// Map route to handler functions
	// Default Handler
	rtr.HandleFunc("/", handler.WelcomeHandler(lgr)).Methods(http.MethodGet)

	// Routes without session ID, require `cfg` to read data store location
	rtr.HandleFunc("/mess_menu", handler.MessMenuHandler(lgr, cfg)).Methods(http.MethodGet)
	rtr.HandleFunc("/contacts", handler.ContactsHandler(lgr, cfg)).Methods(http.MethodGet)

	// Should be the last to allow sieving of requests
	rtr.HandleFunc("/{*}", handler.FourOhFourHandler(lgr)).Methods(http.MethodGet)

	return rtr
}
