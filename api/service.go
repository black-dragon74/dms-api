package api

import (
	"github.com/black-dragon74/dms-api/config"
)

type DMSService struct {
	Session Session
	cfg     *config.Config
}

func NewDMSService(sid string, config *config.Config) DMSService {
	return DMSService{
		Session: newSession(sid),
		cfg:     config,
	}
}
