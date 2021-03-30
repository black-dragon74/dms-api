package handler

import (
	"fmt"
	"github.com/black-dragon74/dms-api/types"
	"go.uber.org/zap"
	"net/http"
)

func MessMenuHandler(lgr *zap.Logger, store *types.GlobalDataStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [MessMenuHandler] Handling /mess_menu")

		_, err := writer.Write(store.MessMenuData)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [MessMenuHandler] [Write] %s", err.Error()))
		}
	}
}
