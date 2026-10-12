package permission

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	permissionhandler "github.com/rajeshbond/smart/internal/http/permission/permission_handler"
	permissionservice "github.com/rajeshbond/smart/internal/http/permission/permission_service"
	permissionstore "github.com/rajeshbond/smart/internal/http/permission/permission_store"
)

type PermissionModule struct {
	PermissionHandler *permissionhandler.PermissionHandler
	PermissionService *permissionservice.PermissionService
	PermissionStore   *permissionstore.PermissionStore
	tokenAuth         *jwtauth.JWTAuth
}

func NewPermissionModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *PermissionModule {
	permissionStore := permissionstore.NewPermissionStore(db)
	permissionService := permissionservice.NewPermissionService(permissionStore)
	permissionHandler := permissionhandler.NewPermissionHandler(*permissionService, tokenAuth)

	return &PermissionModule{
		PermissionHandler: permissionHandler,
		PermissionService: permissionService,
		PermissionStore:   permissionStore,
		tokenAuth:         tokenAuth,
	}
}
