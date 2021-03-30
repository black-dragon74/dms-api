package router

import (
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/app/handler"
	"github.com/black-dragon74/dms-api/app/middleware"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/initialize"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter(lgr *zap.Logger, cfg *config.Config) *mux.Router {
	rtr := mux.NewRouter()
	rtr.Use(middleware.WithContentJSON)

	store, err := initialize.DataStore(lgr, cfg)
	if err != nil {
		lgr.Error(fmt.Sprintf("[Router] [NewRouter] [DataStore] %s", err.Error()))
		return nil
	}

	// Map route to handler functions
	// Default Handler
	rtr.HandleFunc("/", handler.WelcomeHandler(lgr)).Methods(http.MethodGet)

	// Routes without session ID, require `cfg` to read data store location
	rtr.HandleFunc("/mess_menu", handler.MessMenuHandler(lgr, store.MessMenuData)).Methods(http.MethodGet)
	rtr.HandleFunc("/contacts", handler.ContactsHandler(lgr, store.ContactsData)).Methods(http.MethodGet)

	// Some browsers request for favicon even when the content type is set to JSON, handle that
	rtr.HandleFunc("/favicon.ico", handler.FaviconHandler(lgr)).Methods(http.MethodGet)

	if cfg.GetEnv() == config.DevEnv {
		rtr.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			sess := api.NewSession("hnvahilvkw5ld3gq2un12jfk", lgr)
			sess.Validate()
		}).Methods(http.MethodGet)
	}

	// Should be the last to allow sieving of requests
	rtr.HandleFunc("/{*}", handler.FourOhFourHandler(lgr)).Methods(http.MethodGet)

	return rtr
}
