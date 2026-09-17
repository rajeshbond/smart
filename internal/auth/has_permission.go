package auth

func ItHasPermission(
	claims *UserClaims,
	permission Permission,
	required string,
) bool {

	if claims == nil {
		return false
	}

	// XoomGrid SuperAdmin has full access
	if claims.Role == RoleSuperAdmin {
		return true
	}

	switch required {

	case PermissionAdmin:
		return permission.Admin

	case PermissionCreate:
		return permission.Admin || permission.Create

	case PermissionRead:
		return permission.Admin || permission.Read

	case PermissionUpdate:
		return permission.Admin || permission.Update

	case PermissionDelete:
		return permission.Admin || permission.Delete
	}

	return false
}
