package handler

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func ContactsHandler(lgr *zap.Logger, store []byte) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [ContactsHandler] Handling /contacts")

		_, err := writer.Write(store)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [ContactsHandler] [Write] %s", err.Error()))
		}
	}
}
