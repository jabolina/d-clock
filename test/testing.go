package test

import "time"

var ConcurrencyLevel = 5000

func WaitOrTimeout(cb func(), duration time.Duration) bool {
	done := make(chan bool)
	go func() {
		cb()
		done <- true
	}()
	select {
	case <-done:
		return true
	case <-time.After(duration):
		return false
	}
}
