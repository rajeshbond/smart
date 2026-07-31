package shift

import (
	"fmt"
	"strings"
)

func toMinutes(t string) (int, error) {
	var h, m int

	_, err := fmt.Sscanf(t, "%d:%d", &h, &m)
	if err != nil {
		return 0, fmt.Errorf("invalid time format: %s", t)
	}

	if h < 0 || h > 23 || m < 0 || m > 59 {
		return 0, fmt.Errorf("invalid time value: %s", t)
	}

	return h*60 + m, nil
}

type ShiftInterval struct {
	ShiftName string
	Start     int // Minutes from midnight
	End       int // Minutes from midnight (overnight adjusted)
	Duration  int // Duration in minutes
}

func BuildShiftInterval(
	shiftName string,
	shiftStart string,
	shiftEnd string,
) (*ShiftInterval, error) {

	start, err := toMinutes(shiftStart)
	if err != nil {
		return nil, err
	}

	end, err := toMinutes(shiftEnd)
	if err != nil {
		return nil, err
	}

	// Overnight shift
	if end <= start {
		end += 1440
	}

	return &ShiftInterval{
		ShiftName: strings.TrimSpace(shiftName),
		Start:     start,
		End:       end,
		Duration:  end - start,
	}, nil
}
