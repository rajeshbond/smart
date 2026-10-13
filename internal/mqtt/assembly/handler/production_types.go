package handler

import (
	"context"
	"sync"
	"sync/atomic"

	service "github.com/rajeshbond/smart/internal/mqtt/assembly/service"
)

// ============================================================
// PRODUCTION JOB
// ============================================================

type productionJob struct {
	payload []byte
}

// ============================================================
// PRODUCTION HANDLER
// ============================================================

type ProductionHandler struct {

	// --------------------------------------------------------
	// SERVICE
	// --------------------------------------------------------

	service service.ProductionSerive

	// --------------------------------------------------------
	// QUEUE
	// --------------------------------------------------------

	queue chan productionJob

	// --------------------------------------------------------
	// CONTEXT
	// --------------------------------------------------------

	ctx    context.Context
	cancel context.CancelFunc

	// --------------------------------------------------------
	// STATE
	// --------------------------------------------------------

	mu        sync.RWMutex
	accepting bool

	// --------------------------------------------------------
	// WORKERS
	// --------------------------------------------------------

	workersWG sync.WaitGroup

	workerCount atomic.Int32

	// --------------------------------------------------------
	// METRICS
	// --------------------------------------------------------

	receivedCount  atomic.Uint64
	queuedCount    atomic.Uint64
	processedCount atomic.Uint64
	failedCount    atomic.Uint64
	droppedCount   atomic.Uint64
	retryCount     atomic.Uint64

	// --------------------------------------------------------
	// SUPERVISORS
	// --------------------------------------------------------

	supervisorWG sync.WaitGroup
}
