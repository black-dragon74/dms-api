package api

import (
	"github.com/black-dragon74/dms-api/api/internal"
	"github.com/black-dragon74/dms-api/config"
	"github.com/go-redis/redis/v8"
)

type DMSSession struct {
	session internal.Session
	cfg     *config.Config
	rds     *redis.Client
}

func NewDMSSession(sessionID string, cfg *config.Config, rds *redis.Client) DMSSession {
	svc := DMSSession{
		session: internal.NewSession(sessionID, cfg, rds),
		cfg:     cfg,
		rds:     rds,
	}

	return svc
}
