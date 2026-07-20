package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Service) GetProductionLogByTenantIDAndDeviceID(ctx context.Context, req dto.GetProductionRequest) ([]dto.ProductionResponse, error) {
	return s.Store.GetByTenantAndDeviceID(ctx, req)
}
