package permissionstore

import (
	"context"
	"database/sql"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionStore) update(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
	req *permissiondto.UpdatePermissionRequest,
) (*permissiondto.Permission, error) {

	var permission permissiondto.Permission

	err := tx.QueryRowContext(
		ctx,
		queryUpdatePermission,
		req.AllPerm,
		req.CreatePerm,
		req.ReadPerm,
		req.UpdatePerm,
		req.DeletePerm,
		req.UpdatedBy,
		id,
	).Scan(
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
		return nil, err
	}

	return &permission, nil
}
