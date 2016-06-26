package worker

import "sync"

// Request Base type of all work objects
type Request interface{}

// RestoreWork struct contains source, destination and WAL for restore
type RestoreWork struct {
	Source      string
	Destination string
	WalDir      string
}

// BackupWork struct contains source and destination for backup
type BackupWork struct {
	Source      string
	Destination string
}

// Worker for now
type Worker struct {
	Queue  chan Request
	Errs   chan error
	Op     func(Request) error
	Marker *sync.WaitGroup
}

// Start a worker
func (w *Worker) Start() {
	go w.run()
}

func (w *Worker) run() {
	for work := range w.Queue {
		if err := w.Op(work); err != nil {
			w.Errs <- err
		}
		w.Marker.Done()
	}
}
