package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func CacheSession(sid string, rds *redis.Client) error {
	return rds.Set(context.Background(), sid, sid, time.Minute*VarRedisTimeout).Err()
}

func UpdateSessionExpiry(key string, rds *redis.Client) error {
	return rds.Expire(context.Background(), key, time.Minute*VarRedisTimeout).Err()
}
