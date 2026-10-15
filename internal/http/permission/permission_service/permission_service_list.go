package permissionservice

import (
	"context"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionService) List(
	ctx context.Context,
	filter *permissiondto.PermissionFilter,
) ([]*permissiondto.Permission, int, error) {

	permissions, total, err := s.PermissionStore.List(
		ctx,
		filter,
	)

	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}
