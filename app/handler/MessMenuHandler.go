package handler

import (
	"encoding/json"
	"fmt"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func MessMenuHandler(lgr *zap.Logger, store *types.MessMenuModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [MessMenuHandler] Handling /mess_menu")

		data, err := json.Marshal(store)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [MessMenuHandler] [Write] %s", err.Error()))
		}
	}
}
