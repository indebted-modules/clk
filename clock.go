package clk

import (
	"math"
	"time"
)

const dateFormat = "2006-01-02"

// FrozenTime for testing
var FrozenTime = time.Date(1985, 10, 26, 1, 22, 0, 0, time.UTC)

// Clock abstraction
type Clock interface {
	Now() time.Time
}

// SystemClock implementation
type SystemClock struct{}

// FrozenClock implementation
type FrozenClock struct {
	ClockTime time.Time
}

// Now returns system time
func (c *SystemClock) Now() time.Time {
	return time.Now().Truncate(time.Millisecond).UTC()
}

// Now returns frozen time
func (c *FrozenClock) Now() time.Time {
	if c.ClockTime.IsZero() {
		return FrozenTime
	}

	return c.ClockTime
}

// DaysSince returns number of days since given time
func DaysSince(t time.Time) int64 {
	d := time.Now().Sub(t)
	return int64(math.Floor(d.Hours() / 24))
}

// ParseDate converts string formatted as YYYY-MM-DD into UTC Time at midnight
func ParseDate(s string) (time.Time, error) {
	return time.Parse(dateFormat, s)
}

// Unix returns the Time in UTC corresponding to the given Unix time
func Unix(seconds int64) time.Time {
	return time.Unix(seconds, 0).UTC()
}
