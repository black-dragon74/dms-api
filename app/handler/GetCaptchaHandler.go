package handler

import (
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func GetCaptchaHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [GetCaptchaHandler] Handling /captcha")
		resp, err := api.GetCaptcha(lgr)
		if err != nil {
			_, _ = writer.Write(utils.ErrorToJSON(err.Error()))
			return
		}

		_, err = writer.Write(resp)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [GetCaptchaHandler] [Write] %s", err.Error()))
		}
	}
}
