package handler

import (
	"context"
	"log"

	service "github.com/rajeshbond/smart/internal/mqtt/assembly/service"
)

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
	// Start worker supervisor
	// --------------------------------------------------------

	h.supervisorWG.Add(1)

	go h.workerSupervisor()

	// --------------------------------------------------------
	// Start statistics logger
	// --------------------------------------------------------

	h.supervisorWG.Add(1)

	go h.statisticsLogger()

	// --------------------------------------------------------
	// Startup log
	// --------------------------------------------------------

	log.Println("============================================================")
	log.Println("Production Handler Assembly Started")
	log.Printf("Queue Size        : %d", ProductionQueueSize)
	log.Printf("Minimum Workers   : %d", MinProductionWorkers)
	log.Printf("Maximum Workers   : %d", MaxProductionWorkers)
	log.Printf("DB Timeout        : %v", ProductionDBTimeout)
	log.Printf("DB Retries        : %d", ProductionSaveRetries)
	log.Println("============================================================")

	return h
}
