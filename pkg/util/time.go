package util

import (
	"time"
)

// Timestamp returns the curent time in seconds UTC
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampNano returns the curent time in miliseconds UTC
func TimestampNano() int64 {
	return time.Now().UnixNano()
}

// IncT increments a timstamp (in seconds) by m minutes.
func IncT(t int64, m int) int64 {
	return t + (int64)(m*60)
}

// ElapsedTimeSince returns the difference between t and now.
func ElapsedTimeSince(t time.Time) int64 {
	d := time.Since(t)
	return (int64)(d / time.Millisecond)
}

// TimestampToUTC converts a timestamp to UTC timestamp
func TimestampToUTC(t int64) string {
	return time.Unix(t, 0).UTC().String()
}

// TimestampToWeekday retuns the day of the week for the timestamp
func TimestampToWeekday(t int64) int {
	return int(time.Unix(t, 0).Weekday())
}

// TimestampToHour retuns the hour of the day for the timestamp
func TimestampToHour(t int64) int {
	return time.Unix(t, 0).Hour()
}

// StringToTime converts a string with a date/time into a timestamp
func StringToTime(format, t string) int64 {
	tt, err := time.Parse(format, t)
	if err != nil {
		return 0
	}
	return tt.Unix()
}
