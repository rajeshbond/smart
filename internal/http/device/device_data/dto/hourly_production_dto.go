package dto

import "time"

// ============================================================
// REQUEST
// ============================================================
//
// tenant_id is NOT supplied by frontend.
// It will be populated from JWT claims by Handler.
//
// ============================================================

type SaveHourlyProductionRequest struct {
	TenantID   string  `json:"tenant_id"`
	CustomerID *string `json:"customer_id"`
	DeviceID   string  `json:"device_id"`
	MachineID  string  `json:"machine_id"`
	Station    string  `json:"station"`
	Variant    *string `json:"variant"`
}

// ============================================================
// CURRENT SHIFT
// ============================================================

type ShiftInfoHR struct {
	TenantShiftID int64

	ShiftTimingID int64

	ShiftName string

	ShiftStart time.Time

	ShiftEnd time.Time
}

// ============================================================
// HOURLY SLOT
// ============================================================

type HourSlot struct {
	ShiftHourSlotID int64

	ShiftTimingID int64

	SlotIndex int

	SlotStart time.Time

	SlotEnd time.Time
}

// ============================================================
// PRODUCTION STATISTICS
// ============================================================

type HourlyProductionStats struct {
	StartProductionCount int64

	EndProductionCount int64

	HourlyProductionCount int64

	CycleCount int

	MinCycleTimeSec float64

	AvgCycleTimeSec float64

	MaxCycleTimeSec float64
}

// ============================================================
// HOURLY PRODUCTION
// ============================================================

type HourlyProduction struct {
	ID int64 `json:"id"`

	// --------------------------------------------------------
	// Tenant / Device
	// --------------------------------------------------------

	TenantID string `json:"tenant_id"`

	CustomerID string `json:"customer_id,omitempty"`

	DeviceID string `json:"device_id"`

	MachineID string `json:"machine_id"`

	Station string `json:"station"`

	Variant string `json:"variant,omitempty"`

	// --------------------------------------------------------
	// Production Day
	// --------------------------------------------------------

	ProductionDay time.Time `json:"production_day"`

	// --------------------------------------------------------
	// Shift / Slot
	// --------------------------------------------------------

	TenantShiftID int64 `json:"tenant_shift_id"`

	ShiftTimingID int64 `json:"shift_timing_id"`

	ShiftHourSlotID int64 `json:"shift_hour_slot_id"`

	SlotIndex int `json:"slot_index"`

	SlotStart time.Time `json:"slot_start"`

	SlotEnd time.Time `json:"slot_end"`

	// --------------------------------------------------------
	// Actual Timestamp
	// --------------------------------------------------------

	ActualStart time.Time `json:"actual_start"`

	ActualEnd time.Time `json:"actual_end"`

	// --------------------------------------------------------
	// Production
	// --------------------------------------------------------

	StartProductionCount int64 `json:"start_production_count"`

	EndProductionCount int64 `json:"end_production_count"`

	HourlyProductionCount int64 `json:"hourly_production_count"`

	// --------------------------------------------------------
	// Cycle Time
	// --------------------------------------------------------

	CycleCount int `json:"cycle_count"`

	MinCycleTimeSec float64 `json:"min_cycle_time_sec"`

	AvgCycleTimeSec float64 `json:"avg_cycle_time_sec"`

	MaxCycleTimeSec float64 `json:"max_cycle_time_sec"`

	// --------------------------------------------------------
	// Audit
	// --------------------------------------------------------

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}
