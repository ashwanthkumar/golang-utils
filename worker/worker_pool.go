package worker

import (
	"sync"

	"github.com/hashicorp/go-multierror"
)

// Pool is a wrapper to manage a set of Workers efficiently
type Pool struct {
	MaxWorkers  int
	Op          func(Request) error
	workers     []Worker
	items       chan Request
	itemsMarker sync.WaitGroup
	errs        chan error
	finalError  error
}

// Initialize the pool
func (pool *Pool) Initialize() {
	pool.items = make(chan Request)
	pool.errs = make(chan error)
	// Error handler
	go func(combined *error) {
		var result = *combined
		for err := range pool.errs {
			result = multierror.Append(result, err)
		}
		combined = &result
	}(&pool.finalError)
}

// AddWork to a worker in the Pool
func (pool *Pool) AddWork(work Request) {
	if len(pool.workers) < pool.MaxWorkers {
		worker := Worker{
			Queue:  pool.items,
			Errs:   pool.errs,
			Op:     pool.Op,
			Marker: &pool.itemsMarker,
		}
		worker.Start()
		pool.workers = append(pool.workers, worker)
	}
	pool.itemsMarker.Add(1)
	pool.items <- work
}

// Join waits for all the tasks to complete - pool is not usable after this
func (pool *Pool) Join() error {
	close(pool.items)
	pool.itemsMarker.Wait()

	close(pool.errs)
	return pool.finalError
}
