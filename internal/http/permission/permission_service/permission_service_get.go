package permissionservice

import (
	"context"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionService) GetByID(
	ctx context.Context,
	id int64,
) (*permissiondto.Permission, error) {

	permission, err := s.PermissionStore.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	return permission, nil
}
