package internalsetup

import (
	"database/sql"

	"github.com/rajeshbond/smart/internal/http/tenant"
	userrole "github.com/rajeshbond/smart/internal/http/user_role"
	"github.com/rajeshbond/smart/internal/http/users"
)

type Module struct {
	Service *Service
}

func NewModule(db *sql.DB) *Module {

	// Initialize stores
	roleStore := userrole.NewStore(db)
	tenantStore := tenant.NewStore(db)
	userStore := users.NewStore(db)

	// Initialize services
	roleService := userrole.NewService(roleStore)
	tenantService := tenant.NewService(tenantStore)
	userService := users.NewService(userStore, users.RoleProvider(roleService), tenantStore)

	// Initialize setup service
	setupService := NewService(db, tenantService, roleService, userService)

	return &Module{
		Service: setupService,
	}
}
