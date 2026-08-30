package handler

import (
	"log"
	"time"
)

// ============================================================
// WORKER SUPERVISOR
// ============================================================

func (h *ProductionHandler) workerSupervisor() {

	defer h.supervisorWG.Done()

	ticker := time.NewTicker(
		WorkerScaleInterval,
	)

	defer ticker.Stop()

	for {

		select {

		case <-ticker.C:

			h.scaleWorkers()

		case <-h.ctx.Done():

			return
		}
	}
}

// ============================================================
// SCALE WORKERS
// ============================================================

func (h *ProductionHandler) scaleWorkers() {

	queueLength := len(h.queue)

	queueCapacity := cap(h.queue)

	if queueCapacity == 0 {
		return
	}

	workers := int(
		h.workerCount.Load(),
	)

	if workers < MinProductionWorkers {
		workers = MinProductionWorkers
	}

	if queueLength == 0 {
		return
	}

	utilization :=
		float64(queueLength) /
			float64(queueCapacity)

	targetWorkers := workers

	switch {

	case utilization >= 0.75:
		targetWorkers = workers + 4

	case utilization >= 0.50:
		targetWorkers = workers + 3

	case utilization >= 0.25:
		targetWorkers = workers + 2

	case utilization >= 0.10:
		targetWorkers = workers + 1
	}

	if targetWorkers > MaxProductionWorkers {
		targetWorkers = MaxProductionWorkers
	}

	for workers < targetWorkers {

		h.startWorker()

		workers++

		log.Printf(
			"IMM Production Worker Scaling | Workers=%d | Queue=%d/%d",
			workers,
			queueLength,
			queueCapacity,
		)
	}
}
