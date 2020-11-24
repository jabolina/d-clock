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

	// Flag to verify if the clock is was destroyed.
	destroyed int32
}

// Implements the Clock interface.
func (i *InMemoryClock) Tick() error {
	if atomic.CompareAndSwapInt32(&i.destroyed, 0x0, 0x0) {
		atomic.AddUint64(&i.time, 1)
		return nil
	}
	return d_clock.ErrClockDestroyed
}

// Implements the Clock interface.
func (i *InMemoryClock) Tack() (uint64, error) {
	return atomic.LoadUint64(&i.time), nil
}

// Implements the Clock interface.
func (i *InMemoryClock) Leap(to uint64) error {
	if atomic.CompareAndSwapInt32(&i.destroyed, 0x0, 0x0) {
		atomic.StoreUint64(&i.time, to)
		return nil
	}
	return d_clock.ErrClockDestroyed
}

func (i *InMemoryClock) Destroy() error {
	if atomic.CompareAndSwapInt32(&i.destroyed, 0x0, 0x1) {
		return nil
	}
	return d_clock.ErrClockDestroyed
}

// Creates a new InMemoryClock.
func NewInMemoryClock() d_clock.Clock {
	return &InMemoryClock{}
}
