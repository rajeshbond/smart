package store

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/assembly_master/dto"
)

func (s *Store) List(ctx context.Context, tenantID int64) ([]dto.AssemblyMasterResponse, error) {

	rows, err := s.db.QueryContext(ctx, listQuert, tenantID)

	if err != nil {
		return nil, fmt.Errorf("list assembly master: %w", err)
	}

	defer rows.Close()

	result := make([]dto.AssemblyMasterResponse, 0)

	for rows.Next() {

		var item dto.AssemblyMasterResponse

		err := rows.Scan(
			&item.ID,
			&item.TenantID,
			&item.MachineID,
			&item.AssemblyName,
			&item.DeviceID,
			&item.Station,
			&item.Variant,
			&item.HourTargetOutput,
			&item.IsActive,
			&item.IsDeleted,
			&item.CreatedBy,
			&item.UpdatedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan assembly master: %w", err)
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate assembly master: %w", err)
	}
	return result, nil
}
