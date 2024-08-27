package helper

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/go-redis/redis/v8"
)

var cacheExpiration = 15 * time.Minute // Updated to a more reasonable cache expiration

// CacheData caches the data in Redis
func CacheData(ctx context.Context, client *redis.Client, key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, jsonData, cacheExpiration).Err()
}

// GetCachedData retrieves the data from Redis
func GetCachedData(ctx context.Context, client *redis.Client, key string, dest interface{}) error {
	slog.Info("ssssswwwwwwwwwwwwwwwwwww")
	jsonData, err := client.Get(ctx, key).Bytes()
	slog.Info("wwwwwwwwwwwwwww", err)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, dest)
}
