package dto

import "time"

// ============================================================
// CREATE ASSEMBLY MASTER
// ============================================================

type CreateAssemblyMasterRequest struct {
	MachineID        string `json:"machine_id"`
	AssemblyName     string `json:"assembly_name"`
	DeviceID         string `json:"device_id"`
	Station          string `json:"station"`
	Variant          string `json:"variant"`
	HourTargetOutput int    `json:"hour_target_output"`
	// 	CreatedBy        *int   `json:"created_by,omitempty"`
}

// ============================================================
// UPDATE ASSEMBLY MASTER
// ============================================================

type UpdateAssemblyMasterRequest struct {
	MachineID        string `json:"machine_id"`
	AssemblyName     string `json:"assembly_name"`
	DeviceID         string `json:"device_id"`
	Station          string `json:"station"`
	Variant          string `json:"variant"`
	HourTargetOutput int    `json:"hour_target_output"`
	IsActive         bool   `json:"is_active"`
	UpdatedBy        *int   `json:"updated_by,omitempty"`
}

// ============================================================
// RESPONSE
// ============================================================

type AssemblyMasterResponse struct {
	ID               int64     `json:"id"`
	TenantID         int       `json:"tenant_id"`
	MachineID        string    `json:"machine_id"`
	AssemblyName     string    `json:"assembly_name"`
	DeviceID         string    `json:"device_id"`
	Station          string    `json:"station"`
	Variant          string    `json:"variant"`
	HourTargetOutput int       `json:"hour_target_output"`
	IsActive         bool      `json:"is_active"`
	IsDeleted        bool      `json:"is_deleted"`
	CreatedBy        *int      `json:"created_by,omitempty"`
	UpdatedBy        *int      `json:"updated_by,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
