package shift

import "time"

//
// ===============================
// SHIFT DTO
// ===============================
//

type ShiftDTO struct {
	ShiftStart string `json:"shift_start" validate:"required"` // HH:MM
	ShiftEnd   string `json:"shift_end" validate:"required"`   // HH:MM
}

//
// ===============================
// CREATE (Single Shift)
// ===============================
//

type CreateShiftTimingRequest struct {
	TenantShiftID int64  `json:"tenant_shift_id" validate:"required"`
	ShiftStart    string `json:"shift_start" validate:"required"`
	ShiftEnd      string `json:"shift_end" validate:"required"`
}

//
// ===============================
// UPDATE / REPLACE
// ===============================
//

type ReplaceShiftTimingRequest struct {
	TenantShiftID int64  `json:"tenant_shift_id" validate:"required"`
	ShiftStart    string `json:"shift_start" validate:"required"`
	ShiftEnd      string `json:"shift_end" validate:"required"`
}

//
// ===============================
// BULK CREATE
// ===============================
//

type BulkShiftRequest struct {
	TenantCode string `json:"tenant_code" validate:"required"`
	ShiftName  string `json:"shift_name" validate:"required"`
	ShiftStart string `json:"shift_start" validate:"required"`
	ShiftEnd   string `json:"shift_end" validate:"required"`
}

type BulkCreateShiftRequest []BulkShiftRequest

//
// ===============================
// RESPONSE
// ===============================
//

type ShiftTimingResponse struct {
	ID            int64      `json:"id"`
	TenantShiftID int64      `json:"tenant_shift_id"`
	ShiftStart    string     `json:"shift_start"`
	ShiftEnd      string     `json:"shift_end"`
	CreatedBy     *int64     `json:"created_by,omitempty"`
	UpdatedBy     *int64     `json:"updated_by,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

type ShiftInfo struct {
	ShiftID    int64
	ShiftName  string
	ShiftStart string
	ShiftEnd   string
}
