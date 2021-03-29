package handler

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func FourOhFourHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Warn(fmt.Sprintf("[Handler] [404] %s is not a valid route", request.URL.Path))

		respData := `{"error":"404. Route not found"}`

		_, err := writer.Write([]byte(respData))
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [404] [Write] %s", err.Error()))
		}
	}
}
