package app

import (
	"github.com/mirjalilova/api-gateway_blacklist/api"
	"github.com/mirjalilova/api-gateway_blacklist/api/handler"
	_ "github.com/mirjalilova/api-gateway_blacklist/docs"
	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
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

	// // Redis
	// rd := redis.NewClient(&redis.Options{
	// 	Addr:     "redis:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	h := handler.NewHandlerStruct()

	router := api.NewGin(h)
	if err := router.Run(cfg.API_GATEWAY); err != nil {
		slog.Error("Failed to start API Gateway: %v", err)
	}
}