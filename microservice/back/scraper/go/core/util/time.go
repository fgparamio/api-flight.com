package util

import (
	"strconv"
	"time"
)

// Timeable represents an time utility
type Timeable interface {
	//
	UnixMilli(t time.Time) int64
	// Get current timestamp in  millis
	MakeTimestampMilli() int64
	// StrTimestampMilli Get current milliseconds in string
	StrTimestampMilli() string
}

// UnixMilli get int64 value from time.Time
func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

//MakeTimestampMilli  Get gurrent Milliseconds
func MakeTimestampMilli() int64 {
	return UnixMilli(time.Now())
}

// StrTimestampMilli get current milliseconds in string
func StrTimestampMilli() string {
	return strconv.FormatInt(MakeTimestampMilli(), 10)
}
