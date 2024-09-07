package app

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/mirjalilova/api-gateway_blacklist/api"
	"github.com/mirjalilova/api-gateway_blacklist/api/handler"
	_ "github.com/mirjalilova/api-gateway_blacklist/docs"
	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
	"github.com/mirjalilova/api-gateway_blacklist/pkg/minio"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {

	// // Kafka
	// brokers := []string{"kafka_medtrack:29092"}
	// pr, err := prd.NewKafkaProducer(brokers)
	// if err != nil {
	// 	slog.Error("Failed to create Kafka producer:", err)
	// 	return
	// }

	// Redis
	rd := redis.NewClient(&redis.Options{
		Addr:     "redis_blacklist:6379",
		Password: "",
		DB:       0,
	})
	if err := rd.Ping(context.Background()).Err(); err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
		return
	}

	// Minio
	minio, err := minio.MinIOConnect(cfg)
	if err != nil {
		slog.Error("Failed to connect to MinIO: %v", err)
	}

	h := handler.NewHandlerStruct(cfg, rd, minio)

	router := api.NewGin(h)
	if err := router.Run(cfg.API_GATEWAY); err != nil {
		slog.Error("Failed to start API Gateway: %v", err)
	}
}
