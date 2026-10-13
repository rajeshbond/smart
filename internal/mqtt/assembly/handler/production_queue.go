package handler

import "time"

// ============================================================
// ENQUEUE PRODUCTION MESSAGE
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
	// Queue is full
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
