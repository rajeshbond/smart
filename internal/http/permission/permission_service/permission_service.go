package permissionservice

import permissionstore "github.com/rajeshbond/smart/internal/http/permission/permission_store"

type PermissionService struct {
	PermissionStore *permissionstore.PermissionStore
}

func NewPermissionService(store *permissionstore.PermissionStore) *PermissionService {
	return &PermissionService{PermissionStore: store}
}
