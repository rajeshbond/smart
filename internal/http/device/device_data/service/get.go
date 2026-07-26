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
	// Shift Information
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
	// Shift Production
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

	shiftProduction := item.ProductionCount - firstShiftCount

	if shiftProduction < 0 {
		shiftProduction = 0
	}

	item.ProductionCount = shiftProduction

	//-------------------------------------------------
	// Shift OEE
	//-------------------------------------------------

	item.OEE = CalculateLiveOEE(
		shift.ShiftStart,
		item.ProductionTime,
		shiftProduction,
	)

	//-------------------------------------------------
	// Production Day Information
	//-------------------------------------------------

	dayInfo, err := s.shiftProvider.GetProductionDay(
		ctx,
		tenantID,
		item.ProductionTime,
	)
	if err != nil {
		return nil, err
	}

	//-------------------------------------------------
	// Debug Logs
	//-------------------------------------------------

	log.Println("====================================")
	log.Println("Production Time :", item.ProductionTime)
	log.Println("Day Start       :", dayInfo.DayStart)
	log.Println("Day End         :", dayInfo.DayEnd)
	log.Println("====================================")

	//-------------------------------------------------
	// First Day Production Count
	//-------------------------------------------------

	firstDayCount, err := s.Store.GetFirstDayProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		dayInfo.DayStart,
		dayInfo.DayEnd,
	)
	if err != nil {
		return nil, err
	}

	log.Println("Latest Counter    :", shiftProduction)
	log.Println("First Day Counter :", firstDayCount)

	//-------------------------------------------------
	// Day Production
	//-------------------------------------------------

	dayProduction := shiftProduction - firstDayCount

	if dayProduction < 0 {
		dayProduction = 0
	}

	item.DayProduction = dayProduction

	//-------------------------------------------------
	// Day OEE
	//-------------------------------------------------

	item.DayOEE = CalculateDayOEE(
		dayInfo.DayStart,
		item.ProductionTime,
		dayProduction,
	)

	return item, nil
}

// func (s *Service) GetProduction(
// 	ctx context.Context,
// 	req dto.GetProductionRequest,
// 	tenantID int64,
// ) (*dto.ProductionResponse, error) {

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

// 	firstShiftCount, err := s.Store.GetFirstShiftProductionCount(
// 		ctx,
// 		req.TenantID,
// 		req.DeviceID,
// 		req.Station,
// 		shift.ShiftStart,
// 		shift.ShiftEnd,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	//-------------------------------------------------
// 	// Shift Production
// 	//-------------------------------------------------

// 	item.ProductionCount = item.ProductionCount - firstShiftCount

// 	//-------------------------------------------------
// 	// Live OEE
// 	//-------------------------------------------------

// 	log.Printf("Shift Start : %v", shift.ShiftStart)
// 	log.Printf("Shift Count : %d", item.ProductionCount)

// 	item.OEE = CalculateLiveOEE(
// 		shift.ShiftStart,
// 		item.ProductionTime,
// 		item.ProductionCount,
// 	)
// 	// log.Printf("OEE : %+v", item.OEE)

// 	return item, nil
// }

// CalculateDayProduction returns today's production count
// based on the latest machine counter.
func (s *Service) CalculateDayProduction(
	ctx context.Context,
	req dto.GetProductionRequest,
	tenantID int64,
	latest *dto.ProductionResponse,
) (int64, error) {

	//---------------------------------------------------------
	// Production Day Information
	//---------------------------------------------------------

	dayInfo, err := s.shiftProvider.GetProductionDay(
		ctx,
		tenantID,
		latest.ProductionTime,
	)
	if err != nil {
		return 0, err
	}

	//---------------------------------------------------------
	// First Counter of Production Day
	//---------------------------------------------------------

	firstDayCount, err := s.Store.GetFirstDayProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		dayInfo.DayStart,
		dayInfo.DayEnd,
	)
	if err != nil {
		return 0, err
	}

	//---------------------------------------------------------
	// Calculate Day Production
	//---------------------------------------------------------

	dayProduction := latest.ProductionCount - firstDayCount

	if dayProduction < 0 {
		dayProduction = 0
	}

	return dayProduction, nil
}

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
