package auth

const (

	// Dashboard
	PermissionDashboardView = "DASHBOARD_VIEW"

	// Production
	PermissionProductionView = "PRODUCTION_VIEW"
	PermissionCounterReset   = "COUNTER_RESET"
	PermissionLabelUpdate    = "LABEL_UPDATE"
	PermissionJobChange      = "JOB_CHANGE"

	// Reports
	PermissionReportView   = "REPORT_VIEW"
	PermissionReportExport = "REPORT_EXPORT"

	// Device
	PermissionDeviceConfig  = "DEVICE_CONFIG"
	PermissionDeviceRestart = "DEVICE_RESTART"
	PermissionOTAUpdate     = "OTA_UPDATE"

	// Tenant
	PermissionTenantView   = "TENANT_VIEW"
	PermissionTenantCreate = "TENANT_CREATE"
	PermissionTenantUpdate = "TENANT_UPDATE"
	PermissionTenantDelete = "TENANT_DELETE"

	// Users
	PermissionUserView   = "USER_VIEW"
	PermissionUserCreate = "USER_CREATE"
	PermissionUserUpdate = "USER_UPDATE"
	PermissionUserDelete = "USER_DELETE"

	// Roles
	PermissionRoleManage = "ROLE_MANAGE"
)

// const (
// 	RoleSuperAdmin = "superadmin"
// 	RoleXoomAdmin  = "xoomadmin"
// 	RoleXoomUser   = "xoomuser"

// 	RoleDistributorAdmin = "distributoradmin"
// 	RoleDistributorUser  = "distributoruser"

// 	RoleTenantAdmin      = "tenantadmin"
// 	RoleTenantSupervisor = "tenantsupervisor"
// 	RoleTenantOperator   = "tenantoperator"
// )
