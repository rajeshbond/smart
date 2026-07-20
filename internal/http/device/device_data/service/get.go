package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Service) GetProductionLogByTenantIDAndDeviceID(ctx context.Context, tenantID, deviceID string) (*dto.ProductionResponse, error) {
	return s.Store.GetByTenantAndDeviceID(ctx, tenantID, deviceID)
}
