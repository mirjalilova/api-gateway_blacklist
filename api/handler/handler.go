package handler

import (
	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
	"github.com/mirjalilova/api-gateway_blacklist/pkg/clients"
)

type HandlerStruct struct {
	Clients clients.Clients
}

func NewHandlerStruct(cnf *config.Config) *HandlerStruct {
	return &HandlerStruct{
		Clients: *clients.NewClients(*cnf),
	}
}
