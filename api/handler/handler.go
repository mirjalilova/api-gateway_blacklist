package handler

import (
	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
	"github.com/mirjalilova/api-gateway_blacklist/pkg/clients"
	"github.com/go-redis/redis/v8"
)

type HandlerStruct struct {
	Clients clients.Clients
	Redis   *redis.Client
}

func NewHandlerStruct(cnf *config.Config, rd *redis.Client) *HandlerStruct {
	return &HandlerStruct{
		Clients: *clients.NewClients(*cnf),
		Redis: rd,
	}
}
