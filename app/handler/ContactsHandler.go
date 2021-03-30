package handler

import (
	"fmt"
	"github.com/black-dragon74/dms-api/types"
	"go.uber.org/zap"
	"net/http"
)

func ContactsHandler(lgr *zap.Logger, store *types.GlobalDataStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [ContactsHandler] Handling /contacts")

		_, err := writer.Write(store.ContactsData)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [ContactsHandler] [Write] %s", err.Error()))
		}
	}
}
