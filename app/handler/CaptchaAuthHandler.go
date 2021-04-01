package handler

import (
	"encoding/json"
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func CaptchaAuthHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [CaptchaAuthHandler] Handling /captcha_auth")

		// Extract vars from the URL query string
		vars := []string{"sessionid", "username", "password", "captcha"}
		queryVars := utils.ParseArgs(request, &vars)

		// Check if all the requested args are supplied to us by the user
		err := utils.ValidateArgs(&vars, &queryVars)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create a new DMS service and ask it to log us in
		dmsService := api.NewDMSService(queryVars["sessionid"], lgr)
		resp, err := dmsService.Login(queryVars["username"], queryVars["password"], queryVars["captcha"])
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Hanlder] [CaptchaAuthHandler] [Write] %s", err.Error()))
		}
	}
}
