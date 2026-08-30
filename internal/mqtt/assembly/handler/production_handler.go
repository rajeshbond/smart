package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
	service "github.com/rajeshbond/smart/internal/mqtt/assembly/service"
)

// ============================================================
// CONFIGURATION
// ============================================================

const (
	// Bounded MQTT -> database queue.
	ProductionQueueSize = 1000

	// Dynamic worker limits.
	MinProductionWorkers = 2
	MaxProductionWorkers = 20

	// Worker scaling check.
	WorkerScaleInterval = 500 * time.Millisecond

	// Worker exits after being idle for this period,
	// provided the minimum worker count is maintained.
	WorkerIdleTimeout = 30 * time.Second

	// Maximum duration of one database operation.
	ProductionDBTimeout = 5 * time.Second

	// Number of database attempts.
	//
	// IMPORTANT:
	// Save must be idempotent using EventID / UNIQUE constraint
	// before retries are enabled.
	ProductionSaveRetries = 3

	// Delay between retries.
	ProductionRetryDelay = 500 * time.Millisecond

	// Maximum time MQTT callback waits for queue capacity.
	ProductionEnqueueTimeout = 2 * time.Second

	// Statistics interval.
	ProductionStatsInterval = 30 * time.Second
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
// JOB
// ============================================================

type productionJob struct {
	payload []byte
}

// ============================================================
// PRODUCTION HANDLER
// ============================================================

type ProductionHandler struct {
	service service.ProductionSerive

	queue chan productionJob

	ctx    context.Context
	cancel context.CancelFunc

	mu        sync.RWMutex
	accepting bool

	workersWG sync.WaitGroup

	workerCount atomic.Int32

	receivedCount  atomic.Uint64
	queuedCount    atomic.Uint64
	processedCount atomic.Uint64
	failedCount    atomic.Uint64
	droppedCount   atomic.Uint64
	retryCount     atomic.Uint64

	supervisorWG sync.WaitGroup
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewProductionHandler(
	productionService *service.ProductionSerive,
) *ProductionHandler {

	if productionService == nil {
		panic("productionService cannot be nil")
	}

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	h := &ProductionHandler{
		service:   *productionService,
		queue:     make(chan productionJob, ProductionQueueSize),
		ctx:       ctx,
		cancel:    cancel,
		accepting: true,
	}

	// --------------------------------------------------------
	// Start minimum workers
	// --------------------------------------------------------

	for i := 0; i < MinProductionWorkers; i++ {
		h.startWorker()
	}

	// --------------------------------------------------------
	// Worker supervisor
	// --------------------------------------------------------

	h.supervisorWG.Add(1)

	go h.workerSupervisor()

	// --------------------------------------------------------
	// Statistics logger
	// --------------------------------------------------------

	h.supervisorWG.Add(1)

	go h.statisticsLogger()

	log.Println("============================================================")
	log.Println("Production Handler Started")
	log.Printf("Queue Size        : %d", ProductionQueueSize)
	log.Printf("Minimum Workers   : %d", MinProductionWorkers)
	log.Printf("Maximum Workers   : %d", MaxProductionWorkers)
	log.Printf("DB Timeout        : %v", ProductionDBTimeout)
	log.Printf("DB Retries        : %d", ProductionSaveRetries)
	log.Println("============================================================")

	return h
}

// ============================================================
// MQTT MESSAGE HANDLER
// ============================================================

func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {

	return func(client paho.Client, msg paho.Message) {

		if msg == nil {
			log.Println("❌ Production MQTT message is nil")
			return
		}

		h.receivedCount.Add(1)

		// Paho owns the MQTT payload.
		// Copy it before processing asynchronously.
		payload := make([]byte, len(msg.Payload()))

		copy(payload, msg.Payload())

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

	// Fast path.
	select {

	case h.queue <- productionJob{
		payload: payload,
	}:

		h.mu.RUnlock()
		return nil

	default:
	}

	h.mu.RUnlock()

	// Queue is full.
	// Wait only for a bounded amount of time.

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

		// ----------------------------------------------------
		// QUEUED PRODUCTION
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
		// IDLE WORKER
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
				"🔥 Production Worker Panic | Worker=%d | Panic=%v | Stack=%s",
				workerID,
				r,
				stackTrace(),
			)
		}
	}()

	var req dto.ProductionDTO

	if err := json.Unmarshal(data, &req); err != nil {

		h.failedCount.Add(1)

		log.Printf(
			"❌ Production JSON Error | Worker=%d | Error=%v",
			workerID,
			err,
		)

		log.Printf(
			"Raw Payload: %q",
			string(data),
		)

		return
	}

	if req.EventID == "" {

		h.failedCount.Add(1)

		log.Printf(
			"❌ Production Validation Error | Worker=%d | event_id is empty",
			workerID,
		)

		return
	}

	if req.Variant == "" {
		req.Variant = "none"
	}

	log.Printf(
		"📥 Production Processing | "+
			"Worker=%d | "+
			"Tenant=%s | "+
			"EventID=%s | "+
			"Variant=%s | "+
			"Device=%s | "+
			"Station=%s | "+
			"Count=%d | "+
			"Cycle=%.3f",
		workerID,
		req.TenantID,
		req.EventID,
		req.Variant,
		req.DeviceID,
		req.Station,
		req.Count,
		req.CycleTimeSec,
	)

	start := time.Now()

	err := h.saveWithRetry(&req)

	if err != nil {

		h.failedCount.Add(1)

		log.Printf(
			"❌ Production Save FAILED | "+
				"Worker=%d | "+
				"EventID=%s | "+
				"Device=%s | "+
				"Station=%s | "+
				"Duration=%v | "+
				"Error=%v",
			workerID,
			req.EventID,
			req.DeviceID,
			req.Station,
			time.Since(start),
			err,
		)

		return
	}

	h.processedCount.Add(1)

	log.Printf(
		"✅ Production Saved | "+
			"Worker=%d | "+
			"EventID=%s | "+
			"Tenant=%s | "+
			"Device=%s | "+
			"Station=%s | "+
			"Count=%d | "+
			"Cycle=%.3f sec | "+
			"Duration=%v",
		workerID,
		req.EventID,
		req.TenantID,
		req.DeviceID,
		req.Station,
		req.Count,
		req.CycleTimeSec,
		time.Since(start),
	)
}

// ============================================================
// SAVE WITH RETRY
// ============================================================

func (h *ProductionHandler) saveWithRetry(
	req *dto.ProductionDTO,
) error {

	var lastErr error

	for attempt := 1; attempt <= ProductionSaveRetries; attempt++ {

		ctx, cancel := context.WithTimeout(
			context.Background(),
			ProductionDBTimeout,
		)

		err := h.service.Save(
			ctx,
			req,
		)

		cancel()

		if err == nil {
			return nil
		}

		lastErr = err

		if attempt == ProductionSaveRetries {
			break
		}

		h.retryCount.Add(1)

		log.Printf(
			"⚠️ Production DB Retry | EventID=%s | Attempt=%d/%d | Error=%v",
			req.EventID,
			attempt,
			ProductionSaveRetries,
			err,
		)

		delay :=
			ProductionRetryDelay *
				time.Duration(attempt)

		timer := time.NewTimer(delay)

		select {

		case <-timer.C:

		case <-h.ctx.Done():

			timer.Stop()

			return fmt.Errorf(
				"handler shutdown during retry: %w",
				h.ctx.Err(),
			)
		}
	}

	return fmt.Errorf(
		"database save failed after %d attempts: %w",
		ProductionSaveRetries,
		lastErr,
	)
}

// ============================================================
// STATISTICS LOGGER
// ============================================================

func (h *ProductionHandler) statisticsLogger() {

	defer h.supervisorWG.Done()

	ticker := time.NewTicker(
		ProductionStatsInterval,
	)

	defer ticker.Stop()

	for {

		select {

		case <-ticker.C:

			h.logStatistics()

		case <-h.ctx.Done():

			return
		}
	}
}

// ============================================================
// STATISTICS
// ============================================================

type ProductionStats struct {
	QueueLength int
	QueueSize   int

	Workers int

	Received  uint64
	Queued    uint64
	Processed uint64
	Failed    uint64
	Dropped   uint64
	Retries   uint64
}

func (h *ProductionHandler) Stats() ProductionStats {

	return ProductionStats{
		QueueLength: len(h.queue),
		QueueSize:   cap(h.queue),

		Workers: int(
			h.workerCount.Load(),
		),

		Received: h.receivedCount.Load(),

		Queued: h.queuedCount.Load(),

		Processed: h.processedCount.Load(),

		Failed: h.failedCount.Load(),

		Dropped: h.droppedCount.Load(),

		Retries: h.retryCount.Load(),
	}
}

func (h *ProductionHandler) logStatistics() {

	stats := h.Stats()

	log.Printf(
		"📊 Production Handler Stats | "+
			"Queue=%d/%d | "+
			"Workers=%d | "+
			"Received=%d | "+
			"Queued=%d | "+
			"Processed=%d | "+
			"Failed=%d | "+
			"Dropped=%d | "+
			"Retries=%d",
		stats.QueueLength,
		stats.QueueSize,
		stats.Workers,
		stats.Received,
		stats.Queued,
		stats.Processed,
		stats.Failed,
		stats.Dropped,
		stats.Retries,
	)
}

// ============================================================
// CLOSE
// ============================================================

func (h *ProductionHandler) Close() {

	h.mu.Lock()

	if !h.accepting {

		h.mu.Unlock()
		return
	}

	h.accepting = false

	h.mu.Unlock()

	log.Println("============================================================")
	log.Println("Stopping Production Handler...")
	log.Printf(
		"Queue before shutdown: %d/%d",
		len(h.queue),
		cap(h.queue),
	)

	// Stop supervisor/statistics.
	h.cancel()

	// Close queue.
	//
	// Workers drain already queued messages before returning.
	close(h.queue)

	// Wait for all workers.
	h.workersWG.Wait()

	// Wait for supervisor goroutines.
	h.supervisorWG.Wait()

	h.logStatistics()

	log.Println("Production Handler stopped successfully")
	log.Println("============================================================")
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

	return string(buf[:n])
}
