package store

import (
	"context"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

type ProductionStore interface {
	Save(ctx context.Context, req *dto.ProductionDTO) error
	// SaveHeartbeat(ctx context.Context, req *dto.HeartbeatDTO) error
	// SaveStatus(ctx context.Context, req *dto.StatusDTO) error
	// SaveAlarm(ctx context.Context, req *dto.AlarmDTO) error
}
