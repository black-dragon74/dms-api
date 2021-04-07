package initialize

import (
	"context"
	"fmt"
	"github.com/black-dragon74/dms-api/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func RedisStore(cfg *config.Config, lgr *zap.Logger) *redis.Client {
	if !cfg.API.UseRedis() {
		lgr.Info("[Initialize] [RedisStore] Skipping redis store init due to config")
		return nil
	}

	lgr.Info("[Initialize] [RedisStore] Connecting to redis store")
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.GetAddress(),
		Password: cfg.Redis.GetPassword(),
		DB:       cfg.Redis.GetDBId(),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		lgr.Error(fmt.Sprintf("[Initialize] [RedisStore] %s", err.Error()))
		panic("Unable to connect to redis backend. Application halted.")
	}

	lgr.Info("[Initialize] [RedisStore] Successfully connected to redis backend")

	return client
}
