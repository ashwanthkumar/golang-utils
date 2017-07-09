// Package sync contains wrappers around golang's sync wrappers
// Inspired from https://groups.google.com/forum/#!msg/golang-nuts/QCQkT_sihWY/jMW_pZmmAAAJ
package sync

import (
	"sync"
	"sync/atomic"
)

// CountWG is a counted WaitGroup useful to get the number of items on which we're still waiting for
// it internally uses sync.WaitGroup and also an int64 counter which is incremented / decremented using
// sync/atomic package so they're safe for concurrent access
type CountWG struct {
	nocopy  noCopy
	wg      sync.WaitGroup
	counter int64
}

// NewCountWG gives a new version of CountWG to be used
func NewCountWG() *CountWG {
	return &CountWG{
		counter: 0,
	}
}

// Add adds delta, which may be negative to the CountWG counter. If the counter becomes zero,
// all goroutines blocked on Wait are released. If the counter goes negative, Add panics.
func (wg *CountWG) Add(delta int) {
	d := int64(delta)
	wg.wg.Add(delta)
	swapped := false
	for !swapped {
		swapped = atomic.CompareAndSwapInt64(&wg.counter, atomic.LoadInt64(&wg.counter), atomic.LoadInt64(&wg.counter)+d)
	}
}

// Wait blocks until the CountWG counter is zero.
func (wg *CountWG) Wait() {
	wg.wg.Wait()
}

// Done decrements the CountWG counter.
func (wg *CountWG) Done() {
	wg.Add(-1)
}

// Count returns the counter associated with this CountWG
func (wg *CountWG) Count() int {
	return int(atomic.LoadInt64(&wg.counter))
}
