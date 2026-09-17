package auth

// ============================================================
// ROLE
// ============================================================

const (
	RoleSuperAdmin = "superadmin"
)

// ============================================================
// PERMISSION
// ============================================================

const (
	PermissionAdmin  = "admin"
	PermissionCreate = "create"
	PermissionRead   = "read"
	PermissionUpdate = "update"
	PermissionDelete = "delete"
)

// ============================================================
// PERMISSION STRUCT
// ============================================================

type Permission struct {
	Admin  bool `json:"admin"`
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}
