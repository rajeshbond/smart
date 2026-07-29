package utils

import "time"

// Global cache for the IST location to avoid loading it on every call
var istLocation *time.Location

func init() {
	var err error
	istLocation, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// Fallback to fixed +5:30 offset if the timezone database is missing
		istLocation = time.FixedZone("IST", 5*3600+30*60)
	}
}

// ToIST converts any time.Time into Indian Standard Time (IST).
func ToIST(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	return t.In(istLocation)
}

// ParseAndConvertToIST parses a timestamp string and ensures it's localized to IST.
func ParseAndConvertToIST(timeStr string) (time.Time, error) {
	layouts := []string{
		"2006-01-02 15:04:05-07:00",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
		time.RFC3339,
	}

	var parsedTime time.Time
	var err error

	for _, layout := range layouts {
		parsedTime, err = time.ParseInLocation(layout, timeStr, istLocation)
		if err == nil {
			return parsedTime.In(istLocation), nil
		}
	}

	return time.Time{}, err
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func TimePtr(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}

	v := *t
	return &v
}
