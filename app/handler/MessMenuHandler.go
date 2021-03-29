package handler

import (
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/config"
	"go.uber.org/zap"
	"net/http"
)

func MessMenuHandler(lgr *zap.Logger, cfg *config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [MessMenuHandler] Handling /mess_menu")

		dms := api.NewDMSService("", cfg)
		data, err := dms.GetMessMenu()
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [MessMenuHandler] [GetMessMenu] %s", err.Error()))
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [MessMenuHandler] [Write] %s", err.Error()))
		}
	}
}
