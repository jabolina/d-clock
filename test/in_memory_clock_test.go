package test

import (
	"github.com/jabolina/d-clock/internal"
	"sync"
	"testing"
	"time"
)

func Test_InMemoryClock(t *testing.T) {
	clk := internal.NewInMemoryClock()
	group := sync.WaitGroup{}
	apply := func() {
		defer group.Done()
		_, err := clk.Tick()
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
