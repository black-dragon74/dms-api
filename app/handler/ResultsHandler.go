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

func ResultsHandler(lgr *zap.Logger, cfg *config.Config, rds *redis.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handler] [ResultsHandler] Handling /results")

		// Parse the args
		reqVars := []string{utils.VarSessionID, utils.VarSemester}
		args := utils.ParseArgs(request, &reqVars)
		err := utils.ValidateArgs(&reqVars, &args)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create new DMSSession
		dms := api.NewDMSSession(args[utils.VarSessionID], cfg, rds)
		resp, err := dms.GetResult(args[utils.VarSemester])
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Create json response
		bytes, err := json.Marshal(resp)
		if err != nil {
			utils.WriteJSONError(writer, err)
			return
		}

		// Write the response
		_, err = writer.Write(bytes)
		if err != nil {
			lgr.Error(fmt.Sprintf("[Handler] [ResultsHandler] [Write] %s", err.Error()))
		}
	}
}
