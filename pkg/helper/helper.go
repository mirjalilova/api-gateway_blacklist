package helper

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var cacheExpiration = 15 * time.Minute  // Updated to a more reasonable cache expiration

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
	jsonData, err := client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// Key does not exist
			return nil
		}
		return err
	}
	return json.Unmarshal(jsonData, dest)
}
