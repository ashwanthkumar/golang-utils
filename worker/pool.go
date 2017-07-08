package worker

import (
	"github.com/ashwanthkumar/golang-utils/sync"

	"github.com/hashicorp/go-multierror"
)

// Pool is a wrapper to manage a set of Workers efficiently
type Pool struct {
	MaxWorkers int
	Op         func(Request) error

	workers     []Worker
	items       chan Request
	itemsMarker sync.CountWG
	errs        chan error
	finalError  error
}

// Initialize the pool
func (pool *Pool) Initialize() {
	pool.items = make(chan Request, pool.MaxWorkers)
	pool.errs = make(chan error, pool.MaxWorkers)
	// Error handler
	go func(combined *error) {
		var result = *combined
		for err := range pool.errs {
			result = multierror.Append(result, err)
			combined = &result
		}
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

// Wait similar to Join, but the pool is still usable after this
func (pool *Pool) Wait() error {
	pool.itemsMarker.Wait()
	return pool.finalError
}

// Count returns the sum of Pending() + ActiveCount()
func (pool *Pool) Count() int {
	return pool.Pending() + pool.ActiveCount()
}

// Pending returns the number of items still pending to be processed
func (pool *Pool) Pending() int {
	return len(pool.items)
}

// ActiveCount is the count of the workers who are active and doing work
func (pool *Pool) ActiveCount() int {
	return pool.itemsMarker.Count()
}
