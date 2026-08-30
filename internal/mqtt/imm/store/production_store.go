package store

import (
	"context"
	"fmt"
	"log"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

func (s *ImmProductionStore) Save(
	ctx context.Context,
	req *dto.ProductionDTO,
) error {

	// --------------------------------------------------------
	// Validate store
	// --------------------------------------------------------

	if s == nil {
		return fmt.Errorf("production store is nil")
	}

	if s.db == nil {
		return fmt.Errorf("database connection is nil")
	}

	// --------------------------------------------------------
	// Validate request
	// --------------------------------------------------------

	if req == nil {
		return fmt.Errorf("production request is nil")
	}

	if req.EventID == "" {
		return fmt.Errorf("event_id is required")
	}

	// --------------------------------------------------------
	// SQL
	// --------------------------------------------------------

	const query = `
		INSERT INTO imm_production_log
		(
			event_id,
			tenant_id,
			customer_id,
			device_id,
			machine_id,
			station,
			production_count,
			cycle_time_sec,
			production_time
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
			$9
		)
		ON CONFLICT (event_id) DO NOTHING
	`

	// --------------------------------------------------------
	// Execute
	// --------------------------------------------------------

	result, err := s.db.ExecContext(
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
	)

	if err != nil {
		return fmt.Errorf(
			"failed to save IMM production: %w",
			err,
		)
	}

	// --------------------------------------------------------
	// Check duplicate
	// --------------------------------------------------------

	rowsAffected, err := result.RowsAffected()
	if err == nil && rowsAffected == 0 {

		log.Printf(
			"⚠️ IMM production already exists | EventID=%s",
			req.EventID,
		)

		return nil
	}

	// --------------------------------------------------------
	// Success
	// --------------------------------------------------------

	log.Printf(
		"✅ IMM production saved | Device=%s | Machine=%s | Count=%d | EventID=%s",
		req.DeviceID,
		req.MachineID,
		req.Count,
		req.EventID,
	)

	return nil
}
