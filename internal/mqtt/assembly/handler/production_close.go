package handler

import "log"

// ============================================================
// CLOSE PRODUCTION HANDLER
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

	// --------------------------------------------------------
	// Stop supervisor/statistics
	// --------------------------------------------------------

	h.cancel()

	// --------------------------------------------------------
	// Close queue
	// --------------------------------------------------------

	close(h.queue)

	// --------------------------------------------------------
	// Wait for workers
	// --------------------------------------------------------

	h.workersWG.Wait()

	// --------------------------------------------------------
	// Wait for supervisor goroutines
	// --------------------------------------------------------

	h.supervisorWG.Wait()

	// --------------------------------------------------------
	// Final statistics
	// --------------------------------------------------------

	h.logStatistics()

	log.Println("Production Handler stopped successfully")
	log.Println("============================================================")
}
