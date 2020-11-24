package test

import (
	"github.com/jabolina/d-clock/internal"
	"github.com/jabolina/d-clock/pkg/d_clock"
	"sync"
	"testing"
	"time"
)

func Test_InMemoryClock(t *testing.T) {
	clk := internal.NewInMemoryClock()
	group := sync.WaitGroup{}
	apply := func() {
		defer group.Done()
		err := clk.Tick()
		if err != nil {
			t.Errorf("error ticking clock: %#v", err)
		}
	}

	group.Add(ConcurrencyLevel)
	for i := 0; i < ConcurrencyLevel; i++ {
		go apply()
	}

	if !WaitOrTimeout(group.Wait, time.Second) {
		t.Errorf("failed waiting ticks")
	}

	current, err := clk.Tack()
	if err != nil {
		t.Errorf("error reading value: %#v", err)
	}

	if current != uint64(ConcurrencyLevel) {
		t.Errorf("expected %d, found %d", ConcurrencyLevel, current)
	}
}

func Test_FailAfterDestruction(t *testing.T) {
	clk := internal.NewInMemoryClock()
	for i := 0; i < ConcurrencyLevel; i++ {
		err := clk.Tick()
		if err != nil {
			t.Errorf("failed ticking clock. %#v", err)
		}
	}

	value, err := clk.Tack()
	if err != nil {
		t.Errorf("failed retrieving clock value. %#v", err)
	}

	if value != uint64(ConcurrencyLevel) {
		t.Errorf("expected clock %d. found %d", ConcurrencyLevel, value)
	}

	err = clk.Destroy()
	if err != nil {
		t.Errorf("failed destroying clock. %#v", err)
	}

	err = clk.Tick()
	if err != d_clock.ErrClockDestroyed {
		t.Errorf("expected %#v, found %#v", d_clock.ErrClockDestroyed, err)
	}
}
