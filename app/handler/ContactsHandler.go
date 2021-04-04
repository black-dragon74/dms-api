package handler

import (
	"encoding/json"
	"fmt"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func ContactsHandler(lgr *zap.Logger, store []types.ContactsModel) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [ContactsHandler] Handling /contacts")

		data, err := json.Marshal(store)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [ContactsHandler] [Write] %s", err.Error()))
		}
	}
}
