package shift

import "errors"

var (
	ErrShiftOverlap      = errors.New("shift overlap detected")
	ErrTotalExceed24Hour = errors.New("total shift duration exceeds 24 hours")
	ErrDuplicateShift    = errors.New("shift already exists")
	ErrInvalidTime       = errors.New("invalid shift time")
)
