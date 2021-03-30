package handler

import (
	"go.uber.org/zap"
	"net/http"
)

func FaviconHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [FaviconHandler] Ignored favicon request")
		writer.WriteHeader(404)
	}
}
