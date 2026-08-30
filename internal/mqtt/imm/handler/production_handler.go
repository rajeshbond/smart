package handler

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	service "github.com/rajeshbond/smart/internal/mqtt/imm/service"
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
	service service.ImmService

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
	productionService *service.ImmService,
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
	// Start worker supervisor
	// --------------------------------------------------------

	h.supervisorWG.Add(1)

	go h.workerSupervisor()

	// --------------------------------------------------------
	// Start statistics logger
	// --------------------------------------------------------

	h.supervisorWG.Add(1)

	go h.statisticsLogger()

	log.Println("============================================================")
	log.Println("IMM Production Handler Started")
	log.Printf("Queue Size      : %d", ProductionQueueSize)
	log.Printf("Minimum Workers : %d", MinProductionWorkers)
	log.Printf("Maximum Workers : %d", MaxProductionWorkers)
	log.Printf("DB Timeout      : %v", ProductionDBTimeout)
	log.Printf("DB Retries      : %d", ProductionSaveRetries)
	log.Println("============================================================")

	return h
}
