package workerpool

var workerPool chan chan ITask
var WorkPool = make(chan ITask, 4096) // maximum number of task can push to work chan for workers working together

func Start(numWorker int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	workerPool = make(chan chan ITask, numWorker)

	// Now, create all of our workers.
	for i := 0; i<numWorker; i++ {
		worker := NewWorker(i+1, workerPool)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkPool:
				go func() {
					worker := <-workerPool
					worker <- work
				}()
			}
		}
	}()
}