package permissionstore

import (
	"context"
	"database/sql"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionStore) getByID(
	ctx context.Context,
	executor interface {
		QueryRowContext(
			context.Context,
			string,
			...any,
		) *sql.Row
	},
	id int64,
) (*permissiondto.Permission, error) {

	var permission permissiondto.Permission

	err := executor.QueryRowContext(
		ctx,
		queryGetPermissionByID,
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
