package store

import (
	"context"
	"database/sql"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Store) GetByTenantAndDeviceID(ctx context.Context, tenantID, deviceID string) (*dto.ProductionResponse, error) {

	var resp dto.ProductionResponse

	err := s.db.QueryRowContext(ctx, GetLatestProductionByDevice, tenantID, deviceID).Scan(
		&resp.TenantID,
		&resp.CustomerID,
		&resp.DeviceID,
		&resp.MachineID,
		&resp.Station,
		&resp.ProductionCount,
		&resp.CycleTimeSec,
		&resp.ProductionTime,
		&resp.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &resp, nil

}
