package imm

import (
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type Handler struct {
	Service *Service
}

func NewHandler(ser *Service) *Handler {
	return &Handler{
		Service: ser,
	}
}

// telemetry handler

func (h *Handler) TelemetryHandler() paho.MessageHandler {
	return func(c paho.Client, msg paho.Message) {
		var data IMMTelemetry

		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			log.Panicln("telemetry error:", err)
			return
		}

		_ = h.Service.ProcessTelemerty(data)

	}
}
