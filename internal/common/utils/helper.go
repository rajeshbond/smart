package utils

import "time"

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
