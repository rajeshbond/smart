package handler

import (
	"encoding/json"
	"log"
	"time"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

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

	// --------------------------------------------------------
	// Decode JSON
	// --------------------------------------------------------

	var req dto.ProductionDTO

	if err := json.Unmarshal(
		data,
		&req,
	); err != nil {

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

	// --------------------------------------------------------
	// Validate Event ID
	// --------------------------------------------------------

	if req.EventID == "" {

		h.failedCount.Add(1)

		log.Printf(
			"❌ Production Validation Error | Worker=%d | event_id is empty",
			workerID,
		)

		return
	}

	// --------------------------------------------------------
	// Default Variant
	// --------------------------------------------------------

	if req.Variant == "" {
		req.Variant = "none"
	}

	// --------------------------------------------------------
	// Processing log
	// --------------------------------------------------------

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

	// --------------------------------------------------------
	// Save
	// --------------------------------------------------------

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

	// --------------------------------------------------------
	// Success
	// --------------------------------------------------------

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
