package utils

import "time"

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func Month(year int, month time.Month) time.Time {
	return Date(year, month, 1)
}

func LastMonthDay(t time.Time) time.Time {
	return Date(t.Year(), t.Month()+1, 0)
}

func TruncToMonth(t time.Time) time.Time {
	return Date(t.Year(), t.Month(), 1)
}

func TruncToDay(t time.Time) time.Time {
	return Date(t.Year(), t.Month(), t.Day())
}

func FmtAsMonth(t time.Time) string {
	return t.Format("2006-01")
}
