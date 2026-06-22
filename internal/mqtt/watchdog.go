package mqtt

import (
	"context"
	"log"
	"time"

	"github.com/rajeshbond/smart/database"
)

func StartWatchdog(db *database.DB, timeout time.Duration) {

	ticker := time.NewTicker(10 * time.Second)

	go func() {

		for range ticker.C {

			rows, err := db.PGX.Query(context.Background(), `
				SELECT tenant_id, machine_id, last_seen
				FROM machine_status
			`)

			if err != nil {
				log.Println(err)
				continue
			}

			now := time.Now()

			for rows.Next() {

				var tenantID int64
				var machineID string
				var lastSeen time.Time

				rows.Scan(&tenantID, &machineID, &lastSeen)

				if now.Sub(lastSeen) > timeout {

					_, err := db.PGX.Exec(context.Background(), `
						UPDATE machine_status
						SET status = 'OFFLINE'
						WHERE tenant_id = $1 AND machine_id = $2
					`, tenantID, machineID)

					if err == nil {
						log.Println("⚠️ MACHINE OFFLINE:", machineID)
					}
				}
			}

			rows.Close()
		}
	}()
}
