package worker

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestRequest struct{}

func TestPoolDoesAllWork(t *testing.T) {
	var totalCounter int32
	pool := Pool{
		MaxWorkers: 1,
		Op: func(req Request) error {
			atomic.AddInt32(&totalCounter, 1)
			return nil
		},
	}
	pool.Initialize()
	for counter := 0; counter < 5; counter++ {
		pool.AddWork(TestRequest{})
	}
	err := pool.Join()
	assert.NoError(t, err)
	assert.Equal(t, 5, int(totalCounter))
}

func TestPoolCreatesOnlyMaxWorkers(t *testing.T) {
	pool := Pool{
		MaxWorkers: 1,
		Op: func(req Request) error {
			return nil
		},
	}
	pool.Initialize()
	for counter := 0; counter < 5; counter++ {
		pool.AddWork(TestRequest{})
	}
	err := pool.Join()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pool.workers))
}
