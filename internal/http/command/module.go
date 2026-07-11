package command

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/internal/mqtt"
)

type Module struct {
	Service *CommandService
	Handler *Handler
}

func NewModule(
	client paho.Client,
) *Module {

	publisher := mqtt.NewPublisher(client)

	service := NewCommandService(publisher)

	Handler := NewHandler(service)

	return &Module{
		Service: service,
		Handler: Handler,
	}
}
