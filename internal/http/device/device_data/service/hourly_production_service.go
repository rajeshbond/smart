package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

// ============================================================
// SAVE PREVIOUS HOUR
// ============================================================
//
// Example:
//
// Current IST:
//
// 13-Aug-2026 22:03
//
// Current Shift:
//
// Shift B
// 20:00 -> 08:00
//
// Previous completed slot:
//
// 21:00 -> 22:00
//
// Production day:
//
// 13-Aug-2026
//
// Then:
//
// assembly_production_log
//        ↓
// calculate production
//        ↓
// assembly_hourly_production
//
// ============================================================

func (s *Service) SavePreviousHour(
	ctx context.Context,
	req dto.SaveHourlyProductionRequest,
) (*dto.HourlyProduction, error) {

	// ========================================================
	// VALIDATION
	// ========================================================

	if s == nil || s.Store == nil {
		return nil, fmt.Errorf(
			"hourly production service/store is nil",
		)
	}

	if req.TenantID == "" {
		return nil, fmt.Errorf(
			"tenant_id is required",
		)
	}

	if req.DeviceID == "" {
		return nil, fmt.Errorf(
			"device_id is required",
		)
	}

	if req.MachineID == "" {
		return nil, fmt.Errorf(
			"machine_id is required",
		)
	}

	if req.Station == "" {
		return nil, fmt.Errorf(
			"station is required",
		)
	}

	// ========================================================
	// CURRENT TIME
	// ========================================================
	//
	// Your current production system is using IST.
	//
	// ========================================================

	istLocation, err := time.LoadLocation(
		"Asia/Kolkata",
	)

	if err != nil {
		return nil, fmt.Errorf(
			"load IST timezone: %w",
			err,
		)
	}

	now := time.Now().In(
		istLocation,
	)

	// ========================================================
	// BEGIN TRANSACTION
	// ========================================================

	tx, err := s.Store.BeginTx(
		ctx, nil,
	)

	if err != nil {
		return nil, err
	}

	// --------------------------------------------------------
	// Rollback automatically if anything fails.
	// --------------------------------------------------------

	committed := false

	defer func() {

		if !committed {
			_ = tx.Rollback()
		}

	}()

	// ========================================================
	// CURRENT SHIFT
	// ========================================================

	shift, err := s.Store.GetCurrentShift(
		ctx,
		tx,
		req.TenantID,
		now,
	)

	if err != nil {
		return nil, err
	}

	// ========================================================
	// PREVIOUS HOUR SLOT
	// ========================================================
	slot, err := s.Store.GetPreviousHourSlot(
		ctx,
		tx,
		req.TenantID,
		shift.ShiftTimingID,
		now,
	)

	if err != nil {
		return nil, err
	}

	// ========================================================
	// PRODUCTION DAY
	// ========================================================
	//
	// For a normal shift:
	//
	// 08:00 -> 20:00
	//
	// production day = current date
	//
	// For overnight shift:
	//
	// 20:00 -> 08:00
	//
	// after midnight belongs to previous production day.
	//
	// ========================================================

	productionDay := now

	if isOvernightShift(
		shift,
	) {

		currentClock := now.Hour()*3600 +
			now.Minute()*60 +
			now.Second()

		shiftStartClock := clockSeconds(
			shift.ShiftStart,
		)

		// ----------------------------------------------------
		// After midnight but before shift end.
		//
		// Example:
		//
		// Shift B:
		// 20:00 -> 08:00
		//
		// Current:
		// 02:03
		//
		// Production day:
		// previous date
		// ----------------------------------------------------

		if currentClock < shiftStartClock {

			productionDay = productionDay.AddDate(
				0,
				0,
				-1,
			)
		}
	}

	// Normalize production day to midnight.

	productionDay = time.Date(
		productionDay.Year(),
		productionDay.Month(),
		productionDay.Day(),
		0,
		0,
		0,
		0,
		istLocation,
	)

	// ========================================================
	// ACTUAL START / END
	// ========================================================
	//
	// Convert slot TIME into actual TIMESTAMPTZ.
	//
	// Example:
	//
	// Production day = 13-Aug
	//
	// slot:
	// 21:00 -> 22:00
	//
	// actual:
	// 13-Aug 21:00
	// 13-Aug 22:00
	//
	// For:
	//
	// 23:00 -> 00:00
	//
	// actual:
	// 13-Aug 23:00
	// 14-Aug 00:00
	//
	// ========================================================

	actualStart := combineDateAndTime(
		productionDay,
		slot.SlotStart,
		istLocation,
	)

	actualEnd := combineDateAndTime(
		productionDay,
		slot.SlotEnd,
		istLocation,
	)

	// --------------------------------------------------------
	// Slot crosses midnight
	// --------------------------------------------------------

	if !actualEnd.After(actualStart) {

		actualEnd = actualEnd.AddDate(
			0,
			0,
			1,
		)
	}

	// ========================================================
	// PRODUCTION STATISTICS
	// ========================================================

	stats, err := s.Store.GetProductionStats(
		ctx,
		tx,

		req.TenantID,
		req.DeviceID,
		req.Station,

		actualStart,
		actualEnd,
	)

	if err != nil {
		return nil, err
	}

	// ========================================================
	// CREATE HOURLY PRODUCTION
	// ========================================================

	item := &dto.HourlyProduction{

		TenantID: req.TenantID,

		CustomerID: *req.CustomerID,

		DeviceID: req.DeviceID,

		MachineID: req.MachineID,

		Station: req.Station,

		Variant: *req.Variant,

		// ----------------------------------------------------
		// Production Day
		// ----------------------------------------------------

		ProductionDay: productionDay,

		// ----------------------------------------------------
		// Shift
		// ----------------------------------------------------

		TenantShiftID: shift.TenantShiftID,

		ShiftTimingID: shift.ShiftTimingID,

		// ----------------------------------------------------
		// Slot
		// ----------------------------------------------------

		ShiftHourSlotID: slot.ShiftHourSlotID,

		SlotIndex: slot.SlotIndex,

		SlotStart: slot.SlotStart,

		SlotEnd: slot.SlotEnd,

		// ----------------------------------------------------
		// Actual timestamp
		// ----------------------------------------------------

		ActualStart: actualStart,

		ActualEnd: actualEnd,

		// ----------------------------------------------------
		// Production
		// ----------------------------------------------------

		StartProductionCount: stats.StartProductionCount,

		EndProductionCount: stats.EndProductionCount,

		HourlyProductionCount: stats.HourlyProductionCount,

		// ----------------------------------------------------
		// Cycle time
		// ----------------------------------------------------

		CycleCount: stats.CycleCount,

		MinCycleTimeSec: stats.MinCycleTimeSec,

		AvgCycleTimeSec: stats.AvgCycleTimeSec,

		MaxCycleTimeSec: stats.MaxCycleTimeSec,
	}

	// ========================================================
	// SAVE
	// ========================================================

	err = s.Store.SaveHourlyProduction(
		ctx,
		tx,
		item,
	)

	if err != nil {
		return nil, err
	}

	// ========================================================
	// COMMIT
	// ========================================================

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf(
			"commit hourly production: %w",
			err,
		)
	}

	committed = true

	return item, nil
}

// ============================================================
// CHECK OVERNIGHT SHIFT
// ============================================================

func isOvernightShift(
	shift *dto.ShiftInfoHR,
) bool {

	start := clockSeconds(
		shift.ShiftStart,
	)

	end := clockSeconds(
		shift.ShiftEnd,
	)

	return start > end
}

// ============================================================
// CLOCK TO SECONDS
// ============================================================

func clockSeconds(
	t time.Time,
) int {

	return t.Hour()*3600 +
		t.Minute()*60 +
		t.Second()
}

// ============================================================
// COMBINE DATE + TIME
// ============================================================

func combineDateAndTime(
	date time.Time,
	clock time.Time,
	location *time.Location,
) time.Time {

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),

		clock.Hour(),
		clock.Minute(),
		clock.Second(),
		clock.Nanosecond(),

		location,
	)
}
