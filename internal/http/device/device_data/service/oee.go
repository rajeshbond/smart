package service

import (
	"math"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

// Standard (ideal) cycle time in seconds.
const IdealCycleTimeSec = 30.0

// CalculateLiveOEE calculates live OEE using the shift start time,
// the timestamp of the latest production event, and the current shift production.
//
// Assumptions for Phase 1:
//   - No downtime tracking  → Availability = 100%
//   - No reject tracking    → Quality = 100%
//   - Ideal cycle time      → 30 sec/part
func CalculateLiveOEE(
	startTime time.Time,
	productionTime time.Time,
	shiftProduction int64,
) dto.OEE {

	//-------------------------------------------------
	// Elapsed shift time
	//-------------------------------------------------

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err == nil {
		startTime = startTime.In(loc)
		productionTime = productionTime.In(loc)
	}

	elapsedSeconds := productionTime.Sub(startTime).Seconds()

	// elapsedSeconds := productionTime.Sub(shiftStart).Seconds()

	if elapsedSeconds <= 0 {
		return dto.OEE{}
	}

	//-------------------------------------------------
	// Expected production
	//-------------------------------------------------

	expectedProduction := elapsedSeconds / IdealCycleTimeSec

	if expectedProduction <= 0 {
		return dto.OEE{}
	}

	//-------------------------------------------------
	// Availability (Phase 1)
	//-------------------------------------------------

	availability := 100.0

	//-------------------------------------------------
	// Performance
	//-------------------------------------------------

	performance :=
		(float64(shiftProduction) / expectedProduction) * 100

	// Optional: cap performance at 100%
	if performance > 100 {
		performance = 100
	}

	//-------------------------------------------------
	// Quality (Phase 1)
	//-------------------------------------------------

	quality := 100.0

	//-------------------------------------------------
	// OEE
	//-------------------------------------------------

	oee :=
		(availability / 100.0) *
			(performance / 100.0) *
			(quality / 100.0) *
			100.0

	return dto.OEE{
		Availability: round2(availability),
		Performance:  round2(performance),
		Quality:      round2(quality),
		OEE:          round2(oee),
	}
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}
