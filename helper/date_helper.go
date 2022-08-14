package helper

import (
	"time"
)

func SetDefaultDate(date time.Time, defaultYear int, defaultMonth time.Month, defaultDay int) time.Time {
	if date.Unix() < 0 {
		return time.Date(defaultYear, defaultMonth, defaultDay, 0, 0, 0, 0, time.UTC)
	}
	return date
}
