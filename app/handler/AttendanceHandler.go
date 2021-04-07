package handler

import (
	"encoding/json"
	"fmt"
	"github.com/black-dragon74/dms-api/api"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/utils"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"net/http"
)

func AttendanceHandler(lgr *zap.Logger, cfg *config.Config, rds *redis.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [AttendanceHandler] Handling /attendance")

		// Parse the args
		args := utils.ParseArgs(request, &utils.SliceSessionID)
		err := utils.ValidateArgs(&utils.SliceSessionID, &args)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create a new DMS session and get the attendance
		dms := api.NewDMSSession(args[utils.VarSessionID], cfg, rds)
		data, err := dms.GetAttendance()
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Convert to JSON
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Write the response
		_, err = writer.Write(bytes)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [AttendanceHandler] [Write] %s", err.Error()))
		}
	}
}
