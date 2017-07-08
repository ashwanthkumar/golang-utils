package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test shamelessly copied from https://golang.org/src/sync/waitgroup_test.go with little modifications

func testCountWG(t *testing.T, wg1 *CountWG, wg2 *CountWG) {
	n := 16
	wg1.Add(n)
	assert.Equal(t, 16, wg1.Count())
	wg2.Add(n)
	assert.Equal(t, 16, wg2.Count())
	exited := make(chan bool, n)
	for i := 0; i != n; i++ {
		go func(i int) {
			wg1.Done()
			wg2.Wait()
			exited <- true
		}(i)
	}
	wg1.Wait()
	for i := 0; i != n; i++ {
		select {
		case <-exited:
			t.Fatal("CountWG released group too soon")
		default:
		}
		wg2.Done()
	}
	for i := 0; i != n; i++ {
		<-exited // Will block if barrier fails to unlock someone.
	}
}

func TestCountWG(t *testing.T) {
	wg1 := &CountWG{}
	wg2 := &CountWG{}

	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 8; i++ {
		testCountWG(t, wg1, wg2)
	}
}
