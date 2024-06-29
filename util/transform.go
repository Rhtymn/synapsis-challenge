package util

import "time"

func EmailToName(email string) string {
	idx := 0
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			idx = i
		}
	}
	return email[:idx]
}

func ToLocalTime(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		time.Local,
	)
}
