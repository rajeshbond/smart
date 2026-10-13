package handler

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// ============================================================
// MQTT PRODUCTION MESSAGE HANDLER
// ============================================================

func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {

	return func(
		client paho.Client,
		msg paho.Message,
	) {

		if msg == nil {
			log.Println("❌ Production MQTT message is nil")
			return
		}

		h.receivedCount.Add(1)

		// ----------------------------------------------------
		// Copy MQTT payload
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
		// Add to queue
		// ----------------------------------------------------

		if err := h.enqueue(payload); err != nil {

			h.droppedCount.Add(1)

			log.Printf(
				"❌ Production Queue Error | Error=%v | Queue=%d/%d",
				err,
				len(h.queue),
				cap(h.queue),
			)

			return
		}

		h.queuedCount.Add(1)
	}
}
