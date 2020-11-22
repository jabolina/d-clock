package internal

import (
	"github.com/jabolina/d-clock/pkg/d_clock"
	"sync/atomic"
)

// A very simple in-memory Clock implementation.
// This version is not distributed but is cheaper and can be used
// by a single process.
type InMemoryClock struct {
	// The actual clock value.
	time uint64
}

// Implements the Clock interface.
func (i *InMemoryClock) Tick() (uint64, error) {
	return atomic.AddUint64(&i.time, 1), nil
}

// Implements the Clock interface.
func (i *InMemoryClock) Tack() (uint64, error) {
	return atomic.LoadUint64(&i.time), nil
}

// Implements the Clock interface.
func (i *InMemoryClock) Leap(to uint64) (uint64, error) {
	atomic.StoreUint64(&i.time, to)
	return to, nil
}

// Creates a new InMemoryClock.
func NewInMemoryClock() d_clock.Clock {
	return &InMemoryClock{}
}
