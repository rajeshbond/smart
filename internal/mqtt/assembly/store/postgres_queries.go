// package store
package store

import (
	"context"
	"fmt"
	"log"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

func (s *PostgresProductionStore) Save(
	ctx context.Context,
	req *dto.ProductionDTO,
) error {

	log.Println("==========save===========")
	log.Println(req)
	log.Println("=========================")

	// Validate event_id before inserting
	if req.EventID == "" {
		return fmt.Errorf("event_id is required")
	}

	const query = `
	INSERT INTO assembly_production_log
	(
		event_id,
		tenant_id,
		customer_id,
		device_id,
		machine_id,
		station,
		production_count,
		cycle_time_sec,
		production_time,
		variant
	)
	VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10
	)
	ON CONFLICT (event_id) DO NOTHING;
`

	res, err := s.db.ExecContext(
		ctx,
		query,
		req.EventID,
		req.TenantID,
		req.CustomerID,
		req.DeviceID,
		req.MachineID,
		req.Station,
		req.Count,
		req.CycleTimeSec,
		req.Timestamp,
		req.Variant,
	)
	if err != nil {
		return err
	}

	// Optional: Log when a duplicate event was safely ignored
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		log.Printf("⚠️ Duplicate event ignored | EventID: %s", req.EventID)
	}

	return nil
}

// import (
// 	"context"
// 	"fmt"

// 	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
// )

// func (s *PostgresProductionStore) Save(
// 	ctx context.Context,
// 	req *dto.ProductionDTO,
// ) error {

// 	// Validate event_id before inserting
// 	if req.EventID == "" {
// 		return fmt.Errorf("event_id is required")
// 	}

// 	const query = `
// 		INSERT INTO assembly_production_log
// 		(
// 			event_id,
// 			tenant_id,
// 			customer_id,
// 			device_id,
// 			machine_id,
// 			station,
// 			production_count,
// 			cycle_time_sec,
// 			production_time
// 		)
// 		VALUES
// 		(
// 			$1,$2,$3,$4,$5,$6,$7,$8,$9
// 		)
// 	`

// 	_, err := s.db.ExecContext(
// 		ctx,
// 		query,
// 		req.EventID,
// 		req.TenantID,
// 		req.CustomerID,
// 		req.DeviceID,
// 		req.MachineID,
// 		req.Station,
// 		req.Count,
// 		req.CycleTimeSec,
// 		req.Timestamp,
// 	)

// 	return err
// }
