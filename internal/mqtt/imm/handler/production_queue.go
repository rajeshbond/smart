package handler

import (
	"errors"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// ============================================================
// ERRORS
// ============================================================

var (
	ErrProductionHandlerClosed = errors.New(
		"production handler is closed",
	)

	ErrProductionQueueFull = errors.New(
		"production queue is full",
	)
)

// ============================================================
// MQTT MESSAGE HANDLER
// ============================================================

func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {

	return func(
		client paho.Client,
		msg paho.Message,
	) {

		if msg == nil {
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
		// Queue
		// ----------------------------------------------------

		if err := h.enqueue(payload); err != nil {

			h.droppedCount.Add(1)

			return
		}

		h.queuedCount.Add(1)
	}
}

// ============================================================
// ENQUEUE
// ============================================================

func (h *ProductionHandler) enqueue(
	payload []byte,
) error {

	h.mu.RLock()

	if !h.accepting {
		h.mu.RUnlock()
		return ErrProductionHandlerClosed
	}

	// --------------------------------------------------------
	// Fast path
	// --------------------------------------------------------

	select {

	case h.queue <- productionJob{
		payload: payload,
	}:

		h.mu.RUnlock()

		return nil

	default:
	}

	h.mu.RUnlock()

	// --------------------------------------------------------
	// Queue full - bounded wait
	// --------------------------------------------------------

	timer := time.NewTimer(
		ProductionEnqueueTimeout,
	)

	defer timer.Stop()

	select {

	case h.queue <- productionJob{
		payload: payload,
	}:

		return nil

	case <-timer.C:

		return ErrProductionQueueFull

	case <-h.ctx.Done():

		return ErrProductionHandlerClosed
	}
}
