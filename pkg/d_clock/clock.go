package d_clock

// The generic interface a Clock must have. Using
// this interface will have basic functionalities.
type Clock interface {
	// The clock increases by 1.
	Tick() (uint64, error)

	// Returns the current clock value.
	Tack() (uint64, error)

	// Leap the clock to the given time.
	Leap(uint64) (uint64, error)
}
