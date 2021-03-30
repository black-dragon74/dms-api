package handler

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func MessMenuHandler(lgr *zap.Logger, store []byte) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [MessMenuHandler] Handling /mess_menu")

		_, err := writer.Write(store)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [MessMenuHandler] [Write] %s", err.Error()))
		}
	}
}
