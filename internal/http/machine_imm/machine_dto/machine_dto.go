package machinedto

import "time"

// ============================================================
// IMM RESPONSE
// ============================================================

type IMM struct {
	Sr               int64      `json:"sr"`
	TenantID         int64      `json:"tenant_id"`
	DeptID           int64      `json:"dept_id"`
	MachineName      string     `json:"machine_name"`
	MachineNo        string     `json:"machine_no"`
	MachineMake      *string    `json:"machine_make,omitempty"`
	TieBarDistance   *float64   `json:"tie_bar_distance,omitempty"`
	PlattenSize      *float64   `json:"platten_size,omitempty"`
	LocationRingSize *float64   `json:"location_ring_size,omitempty"`
	CreatedBy        *int64     `json:"created_by,omitempty"`
	UpdatedBy        *int64     `json:"updated_by,omitempty"`
	IsDeleted        bool       `json:"is_deleted"`
	DeletedBy        *int64     `json:"deleted_by,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

// ============================================================
// CREATE REQUEST
// ============================================================

type CreateRequest struct {
	DeptID           int64    `json:"dept_id" validate:"required"`
	MachineName      string   `json:"machine_name" validate:"required,min=1,max=100"`
	MachineNo        string   `json:"machine_no" validate:"required,min=1,max=100"`
	MachineMake      *string  `json:"machine_make,omitempty" validate:"omitempty,max=100"`
	TieBarDistance   *float64 `json:"tie_bar_distance,omitempty"`
	PlattenSize      *float64 `json:"platten_size,omitempty"`
	LocationRingSize *float64 `json:"location_ring_size,omitempty"`
}

// ============================================================
// UPDATE REQUEST
// ============================================================

type UpdateRequest struct {
	DeptID           int64    `json:"dept_id" validate:"required"`
	MachineName      string   `json:"machine_name" validate:"required,min=1,max=100"`
	MachineNo        string   `json:"machine_no" validate:"required,min=1,max=100"`
	MachineMake      *string  `json:"machine_make,omitempty" validate:"omitempty,max=100"`
	TieBarDistance   *float64 `json:"tie_bar_distance,omitempty"`
	PlattenSize      *float64 `json:"platten_size,omitempty"`
	LocationRingSize *float64 `json:"location_ring_size,omitempty"`
}

// ============================================================
// LIST REQUEST
// ============================================================

type ListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
	DeptID int64  `json:"dept_id"`
}

// ============================================================
// PAGINATION
// ============================================================

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// ============================================================
// LIST RESPONSE
// ============================================================

type ListResponse struct {
	Data       []IMM      `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// ============================================================
// GENERIC MESSAGE RESPONSE
// ============================================================

type MessageResponse struct {
	Message string `json:"message"`
}
