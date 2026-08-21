package handler

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
	service "github.com/rajeshbond/smart/internal/mqtt/assembly/service"
)

const (
	ProductionQueueSize = 1000
	ProductionWorkers   = 5
)

type ProductionHandler struct {
	service service.ProductionSerive

	queue chan []byte

	wg sync.WaitGroup
}

// NewProductionHandler creates the production handler
// and starts the worker pool.
func NewProductionHandler(
	productionService *service.ProductionSerive,
) *ProductionHandler {

	h := &ProductionHandler{
		service: *productionService,
		queue:   make(chan []byte, ProductionQueueSize),
	}

	// Start workers
	for i := 1; i <= ProductionWorkers; i++ {

		h.wg.Add(1)

		go h.worker(i)
	}

	log.Printf(
		"✅ Production Worker Pool Started | Workers=%d | Queue=%d",
		ProductionWorkers,
		ProductionQueueSize,
	)

	return h
}

// ============================================================
// MQTT MESSAGE HANDLER
// ============================================================

func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {

	return func(client paho.Client, msg paho.Message) {

		// Make a safe copy of MQTT payload.
		payload := make([]byte, len(msg.Payload()))
		copy(payload, msg.Payload())

		// ====================================================
		// Put message into queue
		// ====================================================

		select {

		case h.queue <- payload:

			// Successfully queued.

		default:

			// Queue is full.
			log.Printf(
				"❌ Production Queue FULL | Queue=%d/%d | Message dropped",
				len(h.queue),
				cap(h.queue),
			)
		}
	}
}

// ============================================================
// WORKER
// ============================================================

func (h *ProductionHandler) worker(workerID int) {

	defer h.wg.Done()

	log.Printf(
		"Production Worker %d started",
		workerID,
	)

	for data := range h.queue {

		h.processProduction(workerID, data)
	}

	log.Printf(
		"Production Worker %d stopped",
		workerID,
	)
}

// ============================================================
// PROCESS PRODUCTION
// ============================================================

func (h *ProductionHandler) processProduction(
	workerID int,
	data []byte,
) {

	// log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	// log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	// log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	// log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	var req dto.ProductionDTO

	// ========================================================
	// JSON
	// ========================================================

	if err := json.Unmarshal(data, &req); err != nil {

		log.Printf(
			"Production JSON Error | Worker=%d | Error=%v",
			workerID,
			err,
		)

		log.Printf(
			"Raw Payload: %q",
			string(data),
		)

		return
	}
	// log
	// ========================================================
	// Validate Event ID
	// ========================================================

	if req.EventID == "" {

		log.Printf(
			"Production Error | Worker=%d | event_id is empty",
			workerID,
		)

		return
	}

	if req.Variant == "" {
		req.Variant = "none"
	}
	// ========================================================
	// Context
	// ========================================================

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	// log.Println("************************REQ***************")
	// log.Println(req)
	// log.Println("******************************************")

	defer cancel()

	// ========================================================
	// Log
	// ========================================================

	log.Printf(
		"Production Processing | Worker=%d |TenantCode = %s|Variant = %s | Device=%s | Station=%s | Count=%d | Cycle=%.2f",
		workerID,
		req.TenantID,
		req.Variant,
		req.DeviceID,
		req.Station,
		req.Count,
		req.CycleTimeSec,
	)

	// ========================================================
	// Database Save
	// ========================================================

	start := time.Now()

	if err := h.service.Save(ctx, &req); err != nil {

		log.Printf(
			"❌ Save Error | Worker=%d | EventID=%s | Duration=%v | Error=%v",
			workerID,
			req.EventID,
			time.Since(start),
			err,
		)

		return
	}

	// ========================================================
	// Success
	// ========================================================

	log.Printf(
		"✅ Production Saved | Worker=%d | EventID=%s | Tenant_ID = %s|Device=%s | Station=%s | Count=%d | Cycle Time=%.2f sec | Duration=%v",
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
// SHUTDOWN
// ============================================================

func (h *ProductionHandler) Close() {

	log.Println(
		"Stopping Production Worker Pool...",
	)

	close(h.queue)

	h.wg.Wait()

	log.Println(
		"✅ Production Worker Pool stopped",
	)
}

// type ProductionHandler struct {
// 	service service.ProductionSerive
// }

// func NewProductionHandler(service *service.ProductionSerive) *ProductionHandler {
// 	return &ProductionHandler{service: *service}
// }

// func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {
// 	return func(client paho.Client, msg paho.Message) {
// 		// Deep-copy payload safely before passing to background goroutine
// 		payload := make([]byte, len(msg.Payload()))
// 		copy(payload, msg.Payload())

// 		go func(data []byte) {
// 			var req dto.ProductionDTO

// 			if err := json.Unmarshal(data, &req); err != nil {
// 				log.Printf("Production JSON Error : %v", err)
// 				// FIXED: Use 'data' instead of 'msg.Payload()' to prevent data race
// 				log.Printf("Raw Payload: %q", string(data))
// 				return
// 			}

// 			ctx, cancel := context.WithTimeout(
// 				context.Background(),
// 				5*time.Second,
// 			)
// 			defer cancel()

// 			log.Println("===============Device 2 ===================")
// 			log.Println("===Device 2 ====>", req)
// 			log.Println("============================================")
// 			// ctx, cancel := context.WithTimeout(
// 			// 	context.Background(),
// 			// 	5*time.Second,
// 			// )
// 			// defer cancel()

// 			if err := h.service.Save(ctx, &req); err != nil {
// 				log.Printf("Save Error : %v", err)
// 				return
// 			}

// 			log.Printf(
// 				"Production Saved | EventID=%s | Device=%s | Station=%s | Count=%d | Cycle Time=%.2f sec",
// 				req.EventID,
// 				req.DeviceID,
// 				req.Station,
// 				req.Count,
// 				req.CycleTimeSec,
// 			)
// 		}(payload)
// 	}
// }

// package handler

// import (
// 	"github.com/rajeshbond/smart/internal/mqtt/assembly/service"
// )

// type ProductionHandler struct {
// 	service service.ProductionSerive
// }

// func NewProductionHandler(service *service.ProductionSerive) *ProductionHandler {
// 	return &ProductionHandler{service: *service}
// }

// func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {

// 	return func(client paho.Client, msg paho.Message) {

// 		payload := append([]byte(nil), msg.Payload()...)

// 		go func(data []byte) {

// 			var req dto.ProductionDTO

// 			if err := json.Unmarshal(data, &req); err != nil {

// 				log.Printf("Production JSON Error : %v", err)
// 				log.Printf("Raw Payload: %q", string(msg.Payload()))

// 				return
// 			}

// 			ctx, cancel := context.WithTimeout(
// 				context.Background(),
// 				5*time.Second,
// 			)

// 			defer cancel()

// 			if err := h.service.Save(ctx, &req); err != nil {

// 				log.Printf("Save Error : %v", err)

// 				return
// 			}
// 			// fmt.Println(req)
// 			log.Printf(
// 				"Production Saved | EventID=%s | Device=%s | Station=%s | Count=%d | Cycle Time=%.2f sec",
// 				req.EventID,
// 				req.DeviceID,
// 				req.Station,
// 				req.Count,
// 				req.CycleTimeSec,
// 			)
// 		}(payload)

// 	}

// }
