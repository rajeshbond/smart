package dto

import "time"

// ============================================================
// CREATE REQUEST
// ============================================================

type CreateDepartmentRequest struct {
	Department string `json:"department" validate:"required,max=100"`
}

// ============================================================
// UPDATE REQUEST
// ============================================================

type UpdateDepartmentRequest struct {
	Department string `json:"department" validate:"required,max=100"`
}

// ============================================================
// RESPONSE
// ============================================================

type DepartmentResponse struct {
	ID         int        `json:"id"`
	Department string     `json:"department"`
	CreatedBy  *int       `json:"created_by,omitempty"`
	UpdatedBy  *int       `json:"updated_by,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	IsDeleted  bool       `json:"is_deleted"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

type GetDepartmentResponse struct {
	ID         int        `json:"id"`
	Department string     `json:"department"`
	CreatedBy  *int       `json:"created_by,omitempty"`
	UpdatedBy  *int       `json:"updated_by,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	IsDeleted  bool       `json:"is_deleted"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

type ListDepartmentResponse struct {
	ID         int        `json:"id"`
	Department string     `json:"department"`
	CreatedBy  *int       `json:"created_by,omitempty"`
	UpdatedBy  *int       `json:"updated_by,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	IsDeleted  bool       `json:"is_deleted"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
