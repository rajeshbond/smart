package utils

import "time"

func ToISTNew(t time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return t
	}

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		loc,
	)
}
