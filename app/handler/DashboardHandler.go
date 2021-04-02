package handler

import (
	"encoding/json"
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

func DashboardHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [DashboardHandler] Handling /dashboard")

		// Extract and validate the vars
		args := utils.ParseArgs(request, &utils.SliceSessionID)
		err := utils.ValidateArgs(&utils.SliceSessionID, &args)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create a new DMS service and ask it to provide us with `Dashboard` data
		dmsService := api.NewDMSService(args[utils.VarSessionID], lgr)
		resp, err := dmsService.GetDashboard()
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Write to response
		data, _ := json.Marshal(resp)
		_, err = writer.Write(data)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [DashboardHandler] [Write] %s", err.Error()))
		}
	}
}
