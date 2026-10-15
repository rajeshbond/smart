package permissionhandler

import (
	"github.com/go-chi/jwtauth/v5"
	permissionservice "github.com/rajeshbond/smart/internal/http/permission/permission_service"
)

type PermissionHandler struct {
	PermissionService *permissionservice.PermissionService
	tokenAuth         *jwtauth.JWTAuth
}

func NewPermissionHandler(permissionService permissionservice.PermissionService, tokenAuth *jwtauth.JWTAuth) *PermissionHandler {
	return &PermissionHandler{
		tokenAuth:         tokenAuth,
		PermissionService: &permissionService,
	}
}
