package shiftprovider

import (
	"context"
	"time"
)

type ShiftInfo struct {
	ShiftID    int64
	ShiftName  string
	ShiftStart time.Time
	ShiftEnd   time.Time
}

type ShiftProvider interface {
	GetShiftByProductionTime(
		ctx context.Context,
		tenantID int64,
		productionTime time.Time,
	) (*ShiftInfo, error)
}
