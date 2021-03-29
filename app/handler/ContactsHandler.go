package handler

import (
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/config"
	"go.uber.org/zap"
	"net/http"
)

func ContactsHandler(lgr *zap.Logger, cfg *config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [ContactsHandler] Handling /contacts")

		dms := api.NewDMSService("", cfg)
		data, err := dms.GetContacts()
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [ContactsHandler] [GetContacts] %s", err.Error()))
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [ContactsHandler] [Write] %s", err.Error()))
		}
	}
}
