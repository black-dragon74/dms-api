package api

import (
	"github.com/black-dragon74/dms-api/api/internal"
	"go.uber.org/zap"
)

type DMSService struct {
	session internal.Session
	lgr     *zap.Logger
}

func NewDMSService(sessionID string, lgr *zap.Logger) DMSService {
	svc := DMSService{
		session: internal.NewSession(sessionID),
		lgr:     lgr,
	}

	return svc
}
