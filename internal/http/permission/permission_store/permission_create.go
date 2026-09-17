package permissionstore

import (
	"context"
	"database/sql"
	"errors"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionStore) CreatePermission(ctx context.Context, tx *sql.Tx, req *permissiondto.CreatePermissionRequest) (*permissiondto.Permission, error) {
	var permission permissiondto.Permission

	err := tx.QueryRowContext(ctx, queryCreatePermission, req.AllPerm, req.CreatePerm, req.ReadPerm, req.UpdatePerm, req.DeletePerm, req.CreatedBy).Scan(
		&permission.ID,
		&permission.AllPerm,
		&permission.CreatePerm,
		&permission.ReadPerm,
		&permission.UpdatePerm,
		&permission.DeletePerm,
		&permission.CreatedBy,
		&permission.UpdatedBy,
		&permission.DeletedBy,
		&permission.CreatedAt,
		&permission.UpdatedAt,
		&permission.DeletedAt,
		&permission.IsDeleted,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return &permission, nil
}
