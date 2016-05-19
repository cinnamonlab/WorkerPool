package workerpool

import "fmt"

type Worker struct {
	Id int
	WorkerChan chan ITask
	WorkerPool chan chan ITask
	QuitChan chan bool
}

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerPool chan chan ITask) Worker {
	// Create, and return the worker.
	worker := Worker{
		Id:          id,
		WorkerChan:        make(chan ITask),
		WorkerPool: workerPool,
		QuitChan:    make(chan bool)}

	return worker
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w Worker) Start() {
	fmt.Println("started worker ",w.Id)
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerPool <- w.WorkerChan
			select {
			case task := <-w.WorkerChan:
				// Receive a work request.
			fmt.Println("task executed by worker ",w.Id)
				task.Execute()
			case <-w.QuitChan:
				// We have been asked to stop.
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}