package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/assembly_master/dto"
)

func (s *Store) Update(ctx context.Context, tenantID, id int64, req dto.UpdateAssemblyMasterRequest) (*dto.AssemblyMasterResponse, error) {

	var result dto.AssemblyMasterResponse

	err := s.db.QueryRowContext(
		ctx,
		updateQuery,
		req.MachineID,
		req.AssemblyName,
		req.DeviceID,
		req.Station,
		req.Variant,
		req.HourTargetOutput,
		req.IsActive,
		req.UpdatedBy,
		id,
		tenantID,
	).Scan(
		&result.ID,
		&result.TenantID,
		&result.MachineID,
		&result.AssemblyName,
		&result.DeviceID,
		&result.Station,
		&result.Variant,
		&result.HourTargetOutput,
		&result.IsActive,
		&result.IsDeleted,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, fmt.Errorf("update assembly master: %w", err)
	}

	return &result, nil

}
