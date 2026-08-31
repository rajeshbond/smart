package handler

import "time"

// ============================================================
// PRODUCTION HANDLER CONFIGURATION
// ============================================================

const (
	// MQTT -> Database queue
	ProductionQueueSize = 1000

	// Worker limits
	MinProductionWorkers = 2
	MaxProductionWorkers = 20

	// Worker scaling
	WorkerScaleInterval = 500 * time.Millisecond

	// Worker idle timeout
	WorkerIdleTimeout = 30 * time.Second

	// Database operation timeout
	ProductionDBTimeout = 5 * time.Second

	// Database retry count
	ProductionSaveRetries = 3

	// Delay between retries
	ProductionRetryDelay = 500 * time.Millisecond

	// Maximum time MQTT callback waits for queue
	ProductionEnqueueTimeout = 2 * time.Second

	// Statistics logging interval
	ProductionStatsInterval = 30 * time.Second
)
