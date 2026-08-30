package handler

import (
	"encoding/json"
	"log"
	"sync/atomic"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// ============================================================
// HEARTBEAT HANDLER
// ============================================================
//
// Heartbeat is intentionally kept separate from production.
//
// Production:
//     ProductionHandler
//
// Heartbeat:
//     HeartbeatHandler
//
// This handler receives heartbeat MQTT messages and validates
// that the payload is valid JSON.
//
// Database persistence for heartbeat should be added through
// the actual heartbeat service once its DTO/service contract
// is available.
// ============================================================

type HeartbeatHandler struct {
	receivedCount atomic.Uint64
	validCount    atomic.Uint64
	failedCount   atomic.Uint64
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewHeartbeatHandler() *HeartbeatHandler {

	h := &HeartbeatHandler{}

	log.Println("============================================================")
	log.Println("Heartbeat Handler Started")
	log.Println("============================================================")

	return h
}

// ============================================================
// MQTT MESSAGE HANDLER
// ============================================================

func (h *HeartbeatHandler) HeartBeatHandler() paho.MessageHandler {

	return func(
		client paho.Client,
		msg paho.Message,
	) {

		if msg == nil {

			h.failedCount.Add(1)

			log.Println(
				"❌ Heartbeat MQTT message is nil",
			)

			return
		}

		h.receivedCount.Add(1)

		// ----------------------------------------------------
		// Copy Paho payload.
		// ----------------------------------------------------

		payload := make(
			[]byte,
			len(msg.Payload()),
		)

		copy(
			payload,
			msg.Payload(),
		)

		// ----------------------------------------------------
		// Validate JSON.
		// ----------------------------------------------------

		var raw json.RawMessage

		if err := json.Unmarshal(
			payload,
			&raw,
		); err != nil {

			h.failedCount.Add(1)

			log.Printf(
				"❌ Heartbeat JSON Error | Error=%v | Payload=%q",
				err,
				string(payload),
			)

			return
		}

		h.validCount.Add(1)

		log.Printf(
			"💓 Heartbeat Received | Payload=%s",
			string(payload),
		)
	}
}

// ============================================================
// STATS
// ============================================================

type HeartbeatStats struct {
	Received uint64
	Valid    uint64
	Failed   uint64
}

func (h *HeartbeatHandler) Stats() HeartbeatStats {

	return HeartbeatStats{
		Received: h.receivedCount.Load(),
		Valid:    h.validCount.Load(),
		Failed:   h.failedCount.Load(),
	}
}

// ============================================================
// CLOSE
// ============================================================

func (h *HeartbeatHandler) Close() {

	stats := h.Stats()

	log.Printf(
		"💓 Heartbeat Handler Stopped | "+
			"Received=%d | Valid=%d | Failed=%d",
		stats.Received,
		stats.Valid,
		stats.Failed,
	)

	// Keep time imported/useful if heartbeat processing is
	// extended later.
	_ = time.Second
}
