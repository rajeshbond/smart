package dto

import "time"

// ============================================================
// SAVE HOURLY PRODUCTION REQUEST
// ============================================================

type SaveHourlyProductionRequest struct {
	TenantID   string  `json:"tenant_id"`
	CustomerID *string `json:"customer_id,omitempty"`
	DeviceID   string  `json:"device_id"`
	MachineID  string  `json:"machine_id"`
	Station    string  `json:"station"`
	Variant    *string `json:"variant,omitempty"`
}

// ============================================================
// SHIFT INFORMATION
// ============================================================

type ShiftInfoHR struct {
	TenantShiftID int64 `json:"tenant_shift_id"`
	ShiftTimingID int64 `json:"shift_timing_id"`

	ShiftName string `json:"shift_name"`

	ShiftStart time.Time `json:"shift_start"`
	ShiftEnd   time.Time `json:"shift_end"`
}

// ============================================================
// HOUR SLOT
// ============================================================

type HourSlot struct {
	ShiftHourSlotID int64 `json:"shift_hour_slot_id"`

	ShiftTimingID int64 `json:"shift_timing_id"`

	SlotIndex int `json:"slot_index"`

	SlotStart time.Time `json:"slot_start"`

	SlotEnd time.Time `json:"slot_end"`
}

// ============================================================
// HOURLY PRODUCTION STATISTICS
// ============================================================
//
// IMPORTANT:
//
// StartProductionCount
//     = previous boundary FIRST count
//
// EndProductionCount
//     = current boundary FIRST count
//
// HourlyProductionCount
//     = End - Start
//
// If there is no previous boundary:
//
//     HourlyProductionCount = 0
//
// ============================================================

type HourlyProductionStats struct {

	// Previous boundary count.
	//
	// NULL means there was no previous boundary.
	StartProductionCount *int `json:"start_production_count"`

	// Current boundary first count.
	//
	// NULL means there was no production record
	// in the current boundary/hour.
	EndProductionCount *int `json:"end_production_count"`

	// Calculated production.
	HourlyProductionCount int `json:"hourly_production_count"`

	// Cycle statistics.
	CycleCount int `json:"cycle_count"`

	MinCycleTimeSec float64 `json:"min_cycle_time_sec"`

	AvgCycleTimeSec float64 `json:"avg_cycle_time_sec"`

	MaxCycleTimeSec float64 `json:"max_cycle_time_sec"`
}

// ============================================================
// HOURLY PRODUCTION
// ============================================================

type HourlyProduction struct {
	ID int64 `json:"id,omitempty"`

	TenantID string `json:"tenant_id"`

	CustomerID string `json:"customer_id,omitempty"`

	DeviceID string `json:"device_id"`

	MachineID string `json:"machine_id"`

	Station string `json:"station"`

	Variant string `json:"variant,omitempty"`

	// ----------------------------------------------------------
	// Production Day
	// ----------------------------------------------------------

	ProductionDay time.Time `json:"production_day"`

	// ----------------------------------------------------------
	// Shift
	// ----------------------------------------------------------

	TenantShiftID int64 `json:"tenant_shift_id"`

	ShiftTimingID int64 `json:"shift_timing_id"`

	// ----------------------------------------------------------
	// Hour Slot
	// ----------------------------------------------------------

	ShiftHourSlotID int64 `json:"shift_hour_slot_id"`

	SlotIndex int `json:"slot_index"`

	SlotStart time.Time `json:"slot_start"`

	SlotEnd time.Time `json:"slot_end"`

	// ----------------------------------------------------------
	// Actual Time
	// ----------------------------------------------------------

	ActualStart time.Time `json:"actual_start"`

	ActualEnd time.Time `json:"actual_end"`

	// ----------------------------------------------------------
	// Production Counter
	// ----------------------------------------------------------

	StartProductionCount int `json:"start_production_count"`

	EndProductionCount int `json:"end_production_count"`

	HourlyProductionCount int `json:"hourly_production_count"`

	// ----------------------------------------------------------
	// Cycle
	// ----------------------------------------------------------

	CycleCount int `json:"cycle_count"`

	MinCycleTimeSec float64 `json:"min_cycle_time_sec"`

	AvgCycleTimeSec float64 `json:"avg_cycle_time_sec"`

	MaxCycleTimeSec float64 `json:"max_cycle_time_sec"`

	// ----------------------------------------------------------
	// Timestamps
	// ----------------------------------------------------------

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
