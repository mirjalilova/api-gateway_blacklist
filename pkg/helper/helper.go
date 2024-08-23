package helper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var cacheExpiration = 1 * time.Microsecond

// Cache the data in Redis
func CacheData(ctx context.Context, client *redis.Client, key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, jsonData, cacheExpiration).Err()
}

// Get the data from Redis
func GetCachedData(ctx context.Context, client *redis.Client, key string, dest interface{}) error {
	jsonData, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, dest)
}
