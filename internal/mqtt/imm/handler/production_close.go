package handler

import "log"

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
	log.Println("Stopping IMM Production Handler...")
	log.Printf(
		"Queue before shutdown: %d/%d",
		len(h.queue),
		cap(h.queue),
	)

	// --------------------------------------------------------
	// Stop supervisors
	// --------------------------------------------------------

	h.cancel()

	// --------------------------------------------------------
	// Close queue
	// --------------------------------------------------------

	close(h.queue)

	// --------------------------------------------------------
	// Wait workers
	// --------------------------------------------------------

	h.workersWG.Wait()

	// --------------------------------------------------------
	// Wait supervisor goroutines
	// --------------------------------------------------------

	h.supervisorWG.Wait()

	// --------------------------------------------------------
	// Final statistics
	// --------------------------------------------------------

	h.logStatistics()

	log.Println(
		"IMM Production Handler stopped successfully",
	)

	log.Println("============================================================")
}
