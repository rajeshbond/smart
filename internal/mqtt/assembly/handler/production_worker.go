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
	}
}

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
		"🚀 Production Worker Started | Worker=%d | Workers=%d | Queue=%d/%d",
		workerID,
		h.workerCount.Load(),
		len(h.queue),
		cap(h.queue),
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
			"🛑 Production Worker Stopped | Worker=%d | Workers=%d",
			workerID,
			h.workerCount.Load(),
		)
	}()

	log.Printf(
		"Production Worker %d running",
		workerID,
	)

	idleTimer := time.NewTimer(
		WorkerIdleTimeout,
	)

	defer idleTimer.Stop()

	for {

		select {

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
