package utils

import "time"

func ToISTNew(t time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// return t
		return t.In(loc)
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

func timeOnlyIST(t time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return t
	}

	return time.Date(
		1,
		1,
		1,
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		loc,
	)
}
