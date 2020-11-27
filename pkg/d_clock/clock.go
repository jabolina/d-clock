package d_clock

import (
	"errors"
)

var (
	ErrClockDestroyed = errors.New("clock is already destroyed")
)

// The generic interface a Clock must have. Using
// this interface will have basic functionalities.
type Clock interface {
	// The clock increases by 1.
	Tick() error

	// Returns the current clock value.
	Tack() (uint64, error)

	// Leap the clock to the given time.
	Leap(uint64) error

	// Destroy the clock, after the Clock destruction every
	// method call will panic
	Destroy() error
}
