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
	"github.com/mirjalilova/api-gateway_blacklist/pkg/genai"
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
		slog.Error("Failed to connect to Redis", err)
		return
	}

	minioClient, err := minio.MinIOConnect(cfg)
	if err != nil {
		slog.Error("Failed to connect to MinIO", err)
		return
	}

	advice, err := genai.ConnectGenai(cfg)
	if err != nil {
		slog.Error("Failed to connect to GenAI", err)
		return
	}

	h := handler.NewHandlerStruct(cfg, rd, minioClient, advice)

	router := api.NewGin(h)
	if err := router.Run(cfg.API_GATEWAY); err != nil {
		slog.Error("Failed to start API Gateway", err)
	}
}
