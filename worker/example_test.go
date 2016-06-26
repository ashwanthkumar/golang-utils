package worker

import (
	"log"
	"sync/atomic"
)

func Example() {
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
	// Poll is not usable after calling Join
	if err := pool.Join(); err != nil {
		log.Fatal(err)
	}
}
