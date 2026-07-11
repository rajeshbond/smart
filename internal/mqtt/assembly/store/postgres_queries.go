package store

import (
	"context"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

func (s *PostgresProductionStore) Save(
	ctx context.Context,
	req *dto.ProductionDTO,
) error {

	const query = `
		INSERT INTO assembly_production_log
		(
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
			$1,$2,$3,$4,$5,$6,$7,$8
		)
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		req.TenantID,
		req.CustomerID,
		req.DeviceID,
		req.MachineID,
		req.Station,
		req.Count,
		req.CycleTimeSec,
		req.Timestamp,
	)

	return err
}
