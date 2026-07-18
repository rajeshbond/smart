package permission

import "strings"

func CanCreateDevice(role string) bool {
	switch strings.ToLower(role) {
	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser:

		return true
	}

	return false
}

func CanDeleteDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin:

		return true
	}

	return false
}

func CanViewDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser,
		RoleDistributorAdmin,
		RoleTenantAdmin:

		return true
	}

	return false
}

func CanUpdateDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser,
		RoleDistributorAdmin,
		RoleTenantAdmin:

		return true
	}

	return false
}
