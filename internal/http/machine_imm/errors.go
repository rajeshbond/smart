package machine

import "errors"

// ============================================================
// ERRORS
// ============================================================

var (
	ErrNotFound = errors.New(
		"imm not found",
	)

	ErrDuplicate = errors.New(
		"machine number already exists for this tenant",
	)

	ErrDepartmentNotFound = errors.New(
		"department not found",
	)
)
