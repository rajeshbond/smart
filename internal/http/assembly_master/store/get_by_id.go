package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/assembly_master/dto"
)

func (s *Store) GetByID(ctx context.Context, tenantID, id int64) (*dto.AssemblyMasterResponse, error) {

	var result dto.AssemblyMasterResponse

	err := s.db.QueryRowContext(
		ctx,
		getByIDQuery,
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
		return nil, fmt.Errorf("get assembly master: %w", err)
	}

	return &result, nil

}
