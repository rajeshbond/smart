package imm

import (
	"context"
	"encoding/json"
	"log"
	"time"

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

// TelemetryHandler captures IMM network streams concurrently
func (h *Handler) TelemetryHandler() paho.MessageHandler {
	return func(c paho.Client, msg paho.Message) {
		// 🌟 FIX 1: Instantly extract raw payload bytes to free the MQTT network thread
		payloadBytes := msg.Payload()

		// 🌟 FIX 2: Spin off an isolated concurrent worker thread
		go func(data []byte) {
			var telemetry IMMTelemetry

			// Parse JSON safely inside the worker thread
			if err := json.Unmarshal(data, &telemetry); err != nil {
				// 🛠️ FIXED: Changed from log.Panicln to log.Printf so bad data doesn't crash the server
				log.Printf("❌ [IMM Handler] Malformed JSON payload received: %v", err)
				return
			}

			// 🌟 FIX 3: Introduce a processing context timeout so database writes can't hang forever
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Pass the data down to your business logic layer safely
			if err := h.Service.ProcessTelemerty(ctx, telemetry); err != nil {
				log.Printf("❌ [IMM Handler] Failed to process telemetry for machine: %v", err)
			}
		}(payloadBytes)
	}
}

// package imm

// import (
// 	"encoding/json"
// 	"log"

// 	paho "github.com/eclipse/paho.mqtt.golang"
// )

// type Handler struct {
// 	Service *Service
// }

// func NewHandler(ser *Service) *Handler {
// 	return &Handler{
// 		Service: ser,
// 	}
// }

// // telemetry handler

// func (h *Handler) TelemetryHandler() paho.MessageHandler {
// 	return func(c paho.Client, msg paho.Message) {
// 		var data IMMTelemetry

// 		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
// 			log.Panicln("telemetry error:", err)
// 			return
// 		}

// 		_ = h.Service.ProcessTelemerty(data)

// 	}
// }
