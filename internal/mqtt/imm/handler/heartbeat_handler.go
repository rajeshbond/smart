package handler

import (
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

// ============================================================
// HEARTBEAT HANDLER
// ============================================================

type HeartbeatHandler struct{}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewHeartbeatHandler() *HeartbeatHandler {
	return &HeartbeatHandler{}
}

// ============================================================
// MQTT HEARTBEAT HANDLER
// ============================================================

func (h *HeartbeatHandler) HeartbeatHandler() paho.MessageHandler {

	return func(
		client paho.Client,
		msg paho.Message,
	) {

		if msg == nil {
			log.Println("❌ IMM Heartbeat message is nil")
			return
		}

		var req dto.HeartbeatDTO

		if err := json.Unmarshal(
			msg.Payload(),
			&req,
		); err != nil {

			log.Printf(
				"❌ IMM Heartbeat JSON Error | Error=%v",
				err,
			)

			return
		}

		log.Printf(
			"💓 IMM Heartbeat | Device=%s | Machine=%s | Tenant=%s",
			req.DeviceID,
			req.MachineID,
			req.TenantID,
		)
	}
}
