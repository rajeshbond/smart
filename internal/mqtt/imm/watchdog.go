package imm

import (
	"context"
	"log"
	"time"
)

func StartWatchdog(store *Store, timeout time.Duration) {
	// Every 10 seconds, check for disconnected machines
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			// Create an execution context for this loop iteration cycle
			ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

			// 🌟 This runs a single batch SQL update query. No need for store.GetMachines() here!
			query := `
				UPDATE machine_status
				SET status = 'OFFLINE'
				WHERE last_seen < $1 AND status != 'OFFLINE'`

			// Calculate the absolute time threshold cut-off point
			threshold := time.Now().Add(-timeout)

			result, err := store.DB.Exec(ctx, query, threshold)
			if err != nil {
				log.Printf("❌ [Watchdog Error] Failed to execute status sweep: %v", err)
				cancel()
				continue
			}

			// Log how many factory nodes were marked offline in this sweep
			rowsAffected := result.RowsAffected()
			if rowsAffected > 0 {
				log.Printf("⚠️  [Watchdog Alert] Swept database and marked %d machine(s) as OFFLINE", rowsAffected)
			}

			cancel() // Clean up context resources
		}
	}()
}
