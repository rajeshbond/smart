package auth

func IsSuperAdmin(role string) bool {

	return role == RoleSuperAdmin

}

func IsDistributor(role string) bool {
	switch role {
	case RoleDistributorAdmin,
		RoleDistributorService,
		RoleDistributorUser:
		return true
	}
	return false
}

func IsTenant(role string) bool {
	switch role {
	case RoleTenantAdmin,
		RoleTenantSupervisor,
		RoleTenantMaintenance,
		RoleTenantOperator,
		RoleTenantViewer:

		return true
	}

	return false
}
