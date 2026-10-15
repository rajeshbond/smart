package permissionservice

import (
	"context"
	"database/sql"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionService) Create(
	ctx context.Context,
	req *permissiondto.CreatePermissionRequest,
) (*permissiondto.Permission, error) {

	var permission *permissiondto.Permission

	err := s.withTransaction(
		ctx,
		func(tx *sql.Tx) error {

			var err error

			permission, err = s.PermissionStore.Create(
				ctx,
				tx,
				req,
			)

			if err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return permission, nil
}
