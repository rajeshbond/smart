package service

import (
	"context"
	"log"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Service) GetProductionLogByTenantIDAndDeviceID(ctx context.Context, req dto.GetProductionRequest) ([]dto.ProductionResponse, error) {
	return s.Store.GetByTenantAndDeviceID(ctx, req)
}

func (s *Service) GetProduction(
	ctx context.Context,
	req dto.GetProductionRequest,
	tenantID int64,
) (*dto.ProductionResponse, error) {

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
	log.Println("Shift Name ----->", shift.ShiftName)
	item.ShiftName = shift.ShiftName

	//-------------------------------------------------
	// First Shift Count
	//-------------------------------------------------

	firstShiftCount, err := s.Store.GetFirstShiftProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		shift.ShiftStart,
		shift.ShiftEnd,
	)

	if err != nil {
		return nil, err
	}

	//-------------------------------------------------
	// Shift Production
	//-------------------------------------------------

	item.ProductionCount = item.ProductionCount - firstShiftCount

	//-------------------------------------------------
	// Live OEE
	//-------------------------------------------------

	log.Printf("Shift Start : %v", shift.ShiftStart)
	log.Printf("Shift Count : %d", item.ProductionCount)

	item.OEE = CalculateLiveOEE(
		shift.ShiftStart,
		item.ProductionTime,
		item.ProductionCount,
	)
	// log.Printf("OEE : %+v", item.OEE)

	return item, nil
}

// CalculateDayProduction returns today's production count
// based on the latest machine counter.
// func (s *Service) CalculateDayProduction(
// 	ctx context.Context,
// 	req dto.GetProductionRequest,
// 	tenantID int64,
// 	latest *dto.ProductionResponse,
// ) (int64, error) {

// 	//---------------------------------------------------------
// 	// Production Day Information
// 	//---------------------------------------------------------

// 	dayInfo, err := s.shiftProvider.GetProductionDay(
// 		ctx,
// 		tenantID,
// 		latest.ProductionTime,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	//---------------------------------------------------------
// 	// First Counter of Production Day
// 	//---------------------------------------------------------

// 	firstDayCount, err := s.Store.GetFirstDayProductionCount(
// 		ctx,
// 		req.TenantID,
// 		req.DeviceID,
// 		req.Station,
// 		dayInfo.DayStart,
// 		dayInfo.DayEnd,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	//---------------------------------------------------------
// 	// Calculate Day Production
// 	//---------------------------------------------------------

// 	dayProduction := latest.ProductionCount - firstDayCount

// 	if dayProduction < 0 {
// 		dayProduction = 0
// 	}

// 	return dayProduction, nil
// }

// func (s *Service) GetProduction(ctx context.Context, req dto.GetProductionRequest, tenantID int64) (*dto.ProductionResponse, error) {
// 	//-------------------------------------------------
// 	// Latest Production
// 	//-------------------------------------------------
// 	item, err := s.Store.GetTenantCountByDevice(ctx, req)

// 	if err != nil {
// 		return nil, err
// 	}
// 	//-------------------------------------------------
// 	// Shift Info
// 	//-------------------------------------------------
// 	shift, err := s.shiftProvider.GetShiftByProductionTime(
// 		ctx,
// 		tenantID,
// 		item.ProductionTime,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	item.ShiftName = shift.ShiftName

// 	//-------------------------------------------------
// 	// First Shift Count
// 	//-------------------------------------------------

// 	firstshiftCount, err := s.Store.GetFirstShiftProductionCount(ctx, req.TenantID, req.DeviceID, req.Station, shift.ShiftStart, shift.ShiftEnd)
// 	if err != nil {
// 		return nil, err
// 	}

// 	item.ProductionCount = item.ProductionCount - firstshiftCount

// 	return item, nil

// }
