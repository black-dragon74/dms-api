package api

import (
	"github.com/black-dragon74/dms-api/api/internal"
	"go.uber.org/zap"
)

type DMSSession struct {
	session internal.Session
	lgr     *zap.Logger
}

func NewDMSSession(sessionID string, lgr *zap.Logger) DMSSession {
	svc := DMSSession{
		session: internal.NewSession(sessionID),
		lgr:     lgr,
	}

	return svc
}
