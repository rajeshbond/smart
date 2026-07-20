package store

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Store) GetByTenantAndDeviceID(ctx context.Context, req dto.GetProductionRequest) ([]dto.ProductionResponse, error) {
	rows, err := s.db.QueryContext(ctx, GetLatestProductionByDevice, req.TenantID, req.DeviceID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]dto.ProductionResponse, 0)

	for rows.Next() {
		var item dto.ProductionResponse

		err := rows.Scan(
			&item.TenantID,
			&item.CustomerID,
			&item.DeviceID,
			&item.MachineID,
			&item.Station,
			&item.ProductionCount,
			&item.CycleTimeSec,
			&item.ProductionTime,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}

// func (s *Store) GetByTenantAndDeviceID(ctx context.Context, tenantID, deviceID string) (*dto.ProductionResponse, error) {

// 	var resp dto.ProductionResponse

// 	err := s.db.QueryRowContext(ctx, GetLatestProductionByDevice, tenantID, deviceID).Scan(
// 		&resp.TenantID,
// 		&resp.CustomerID,
// 		&resp.DeviceID,
// 		&resp.MachineID,
// 		&resp.Station,
// 		&resp.ProductionCount,
// 		&resp.CycleTimeSec,
// 		&resp.ProductionTime,
// 		&resp.CreatedAt,
// 	)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, sql.ErrNoRows
// 		}
// 		return nil, err
// 	}

// 	return &resp, nil

// }
