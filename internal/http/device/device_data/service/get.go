package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Service) GetProductionLogByTenantIDAndDeviceID(ctx context.Context, req dto.GetProductionRequest) ([]dto.ProductionResponse, error) {
	return s.Store.GetByTenantAndDeviceID(ctx, req)
}

func (s *Service) GetProduction(ctx context.Context, req dto.GetProductionRequest, tenantID int64) (*dto.ProductionResponse, error) {
	//-------------------------------------------------
	// Latest Production
	//-------------------------------------------------
	item, err := s.Store.GetTenantCountByDevice(ctx, req)

	if err != nil {
		return nil, err
	}
	//-------------------------------------------------
	// Shift Info
	//-------------------------------------------------
	shift, err := s.shiftProvider.GetShiftByProductionTime(
		ctx,
		tenantID,
		item.ProductionTime,
	)

	if err != nil {
		return nil, err
	}

	item.ShiftName = shift.ShiftName

	//-------------------------------------------------
	// First Shift Count
	//-------------------------------------------------

	firstshiftCount, err := s.Store.GetFirstShiftProductionCount(ctx, req.TenantID, req.DeviceID, req.Station, shift.ShiftStart, shift.ShiftEnd)
	if err != nil {
		return nil, err
	}

	item.ProductionCount = item.ProductionCount - firstshiftCount

	return item, nil

}
