package handler

import (
	"errors"
	"fmt"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func FeatureNotAvailableHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info(fmt.Sprintf("[Handler] [FeatureNotAvailableHandler] Ignored %s request", request.URL.Path))

		utils.WriteJSONError(writer, errors.New("this feature is not available"))
	}
}
