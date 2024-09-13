package handler

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/generative-ai-go/genai"
	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
	"github.com/mirjalilova/api-gateway_blacklist/pkg/clients"
	"github.com/mirjalilova/api-gateway_blacklist/pkg/minio"
)

type HandlerStruct struct {
	Clients clients.Clients
	Redis   *redis.Client
	Minio   *minio.MinIO
	Genai   *genai.ChatSession
}

func NewHandlerStruct(cnf *config.Config, rd *redis.Client, mn *minio.MinIO, gn *genai.ChatSession) *HandlerStruct {
	return &HandlerStruct{
		Clients: *clients.NewClients(*cnf),
		Redis:   rd,
		Minio:   mn,
		Genai:   gn,
	}
}
