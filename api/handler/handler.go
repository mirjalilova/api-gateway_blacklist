package handler

import (
	"github.com/mirjalilova/api-gateway_blacklist/pkg/clients"
)

type HandlerStruct struct {
	Clients clients.Clients
}

func NewHandlerStruct() *HandlerStruct {
	return &HandlerStruct{
		Clients: *clients.NewClients(),
	}
}
