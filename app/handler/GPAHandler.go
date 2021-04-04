package handler

import (
	"encoding/json"
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func GPAHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [GPAHandler] Handling /gpa")

		args := utils.ParseArgs(request, &utils.SliceSessionID)
		err := utils.ValidateArgs(&utils.SliceSessionID, &args)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create a new DMS session
		dms := api.NewDMSSession(args[utils.VarSessionID], lgr)
		resp, err := dms.GetGPA()
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create bytes
		data, err := json.Marshal(resp)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [GPAHandler] [Write] %s", err.Error()))
		}
	}
}
