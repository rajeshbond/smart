package imm

import (
	"context"
	"log"
	"time"
)

func StartWatchdog(
	store *Store,
	timeout time.Duration,
) {

	ticker := time.NewTicker(10 * time.Second)

	go func() {

		defer ticker.Stop()

		for range ticker.C {

			machines, err := store.GetMAchines()
			if err != nil {
				log.Println("watchdog:", err)
				continue
			}

			now := time.Now()

			query := `
				UPDATE machine_status
				SET status='OFFLINE'
				WHERE tenant_id=$1
				AND machine_id=$2
			`

			for _, m := range machines {

				if now.Sub(m.LastSeen) > timeout {

					_, err := store.DB.PGX.Exec(
						context.Background(),
						query,
						m.TenantID,
						m.MachineID,
					)

					if err != nil {
						log.Println("offline update error:", err)
						continue
					}

					log.Printf(
						"⚠️ OFFLINE Tenant=%d Machine=%s",
						m.TenantID,
						m.MachineID,
					)
				}
			}
		}
	}()
}
