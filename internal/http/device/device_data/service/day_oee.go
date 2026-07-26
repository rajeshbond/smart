package service

import (
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

const (
	DayAvailability = 100.0
	DayQuality      = 100.0
)

// CalculateDayOEE calculates the OEE for the current production day.
func CalculateDayOEE(
	dayStart time.Time,
	productionTime time.Time,
	dayProduction int64,
) dto.OEE {

	elapsedSeconds := productionTime.Sub(dayStart).Seconds()

	if elapsedSeconds <= 0 {
		return dto.OEE{}
	}

	expectedProduction := elapsedSeconds / IdealCycleTimeSec

	if expectedProduction <= 0 {
		return dto.OEE{}
	}

	performance :=
		(float64(dayProduction) / expectedProduction) * 100

	if performance > 100 {
		performance = 100
	}

	availability := DayAvailability
	quality := DayQuality

	oee :=
		(availability * performance * quality) / 10000

	return dto.OEE{
		Availability: round2(availability),
		Performance:  round2(performance),
		Quality:      round2(quality),
		OEE:          round2(oee),
	}
}

// func round2(v float64) float64 {
// 	return math.Round(v*100) / 100
// }
