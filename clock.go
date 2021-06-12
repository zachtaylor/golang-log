package log

import "time"

// Clock wraps time.Now() calls
type Clock interface {
	// Now returns the time
	Now() time.Time
}

// ClockFunc is a func type for Clock
type ClockFunc func() time.Time

// Now returns f()
func (f ClockFunc) Now() time.Time { return f() }

// DefaultClock creates a Clock from NewTime
func DefaultClock() Clock { return ClockFunc(time.Now) }
