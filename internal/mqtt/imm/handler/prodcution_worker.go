package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

// ============================================================
// START WORKER
// ============================================================

func (h *ProductionHandler) startWorker() {

	workerID := int(
		h.workerCount.Add(1),
	)

	h.workersWG.Add(1)

	go h.worker(workerID)

	log.Printf(
		"IMM Production Worker Started | Worker=%d | Workers=%d",
		workerID,
		h.workerCount.Load(),
	)
}

// ============================================================
// WORKER
// ============================================================

func (h *ProductionHandler) worker(
	workerID int,
) {

	defer h.workersWG.Done()

	defer func() {

		h.workerCount.Add(-1)

		log.Printf(
			"IMM Production Worker Stopped | Worker=%d | Workers=%d",
			workerID,
			h.workerCount.Load(),
		)
	}()

	idleTimer := time.NewTimer(
		WorkerIdleTimeout,
	)

	defer idleTimer.Stop()

	for {

		select {

		// ----------------------------------------------------
		// JOB
		// ----------------------------------------------------

		case job, ok := <-h.queue:

			if !ok {
				return
			}

			if !idleTimer.Stop() {

				select {
				case <-idleTimer.C:
				default:
				}
			}

			idleTimer.Reset(
				WorkerIdleTimeout,
			)

			h.processProduction(
				workerID,
				job.payload,
			)

		// ----------------------------------------------------
		// IDLE
		// ----------------------------------------------------

		case <-idleTimer.C:

			if h.workerCount.Load() >
				int32(MinProductionWorkers) {

				return
			}

			idleTimer.Reset(
				WorkerIdleTimeout,
			)
		}
	}
}

// ============================================================
// PROCESS PRODUCTION
// ============================================================

func (h *ProductionHandler) processProduction(
	workerID int,
	data []byte,
) {

	defer func() {

		if r := recover(); r != nil {

			h.failedCount.Add(1)

			log.Printf(
				"IMM Production Worker Panic | Worker=%d | Panic=%v",
				workerID,
				r,
			)

			log.Println(stackTrace())
		}
	}()

	var req dto.ProductionDTO

	// --------------------------------------------------------
	// JSON
	// --------------------------------------------------------

	if err := json.Unmarshal(data, &req); err != nil {

		h.failedCount.Add(1)

		log.Printf(
			"IMM Production JSON Error | Worker=%d | Error=%v",
			workerID,
			err,
		)

		return
	}

	// --------------------------------------------------------
	// Validation
	// --------------------------------------------------------

	if req.EventID == "" {

		h.failedCount.Add(1)

		log.Printf(
			"IMM Production Validation Error | Worker=%d | event_id empty",
			workerID,
		)

		return
	}

	// --------------------------------------------------------
	// Default variant
	// --------------------------------------------------------

	// if req.Variant == "" {
	// 	req.Variant = "none"
	// }

	start := time.Now()

	// --------------------------------------------------------
	// Save
	// --------------------------------------------------------

	if err := h.saveWithRetry(&req); err != nil {

		h.failedCount.Add(1)

		log.Printf(
			"IMM Production Save FAILED | Worker=%d | EventID=%s | Error=%v",
			workerID,
			req.EventID,
			err,
		)

		return
	}

	h.processedCount.Add(1)

	log.Printf(
		"IMM Production Saved | Worker=%d | EventID=%s | Device=%s | Machine=%s | Duration=%v",
		workerID,
		req.EventID,
		req.DeviceID,
		req.MachineID,
		time.Since(start),
	)
}

// ============================================================
// STACK TRACE
// ============================================================

func stackTrace() string {

	buf := make(
		[]byte,
		64*1024,
	)

	n := runtime.Stack(
		buf,
		false,
	)

	return fmt.Sprintf(
		"%s",
		buf[:n],
	)
}
