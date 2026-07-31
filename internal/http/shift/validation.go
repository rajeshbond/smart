package shift

import "fmt"

func ValidateOverlap(
	existing []ShiftInfo,
	newShift *ShiftInterval,
) error {

	for _, s := range existing {

		existingShift, err := BuildShiftInterval(
			s.ShiftName,
			s.ShiftStart,
			s.ShiftEnd,
		)

		if err != nil {
			return err
		}

		// Normal comparison
		if newShift.Start < existingShift.End &&
			newShift.End > existingShift.Start {

			return fmt.Errorf(
				"shift overlap detected between '%s' and '%s'",
				newShift.ShiftName,
				existingShift.ShiftName,
			)
		}

		// ---------
		// Overnight comparison
		// ---------

		if newShift.Start+1440 < existingShift.End &&
			newShift.End+1440 > existingShift.Start {

			return fmt.Errorf(
				"shift overlap detected between '%s' and '%s'",
				newShift.ShiftName,
				existingShift.ShiftName,
			)
		}

		if newShift.Start < existingShift.End+1440 &&
			newShift.End > existingShift.Start+1440 {

			return fmt.Errorf(
				"shift overlap detected between '%s' and '%s'",
				newShift.ShiftName,
				existingShift.ShiftName,
			)
		}
	}

	return nil
}

func ValidateTotalDuration(
	existing []ShiftInfo,
	newShift *ShiftInterval,
) error {

	total := newShift.Duration

	for _, s := range existing {

		existingShift, err := BuildShiftInterval(
			s.ShiftName,
			s.ShiftStart,
			s.ShiftEnd,
		)

		if err != nil {
			return err
		}

		total += existingShift.Duration
	}

	if total > 1440 {

		return fmt.Errorf(
			"total shift duration exceeds 24 hours",
		)
	}

	return nil
}
