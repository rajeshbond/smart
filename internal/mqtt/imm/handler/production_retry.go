package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

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
			"IMM Production DB Retry | EventID=%s | Attempt=%d/%d | Error=%v",
			req.EventID,
			attempt,
			ProductionSaveRetries,
			err,
		)

		delay := ProductionRetryDelay *
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
