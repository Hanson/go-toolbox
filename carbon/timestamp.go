package carbon

import "time"

func StartOfDay() time.Time {
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

func StartOfDayWithTime(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Local")
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}
