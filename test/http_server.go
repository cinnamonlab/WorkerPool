package main

import (
	"fmt"
	"net/http"
	"github.com/cinnamonlab/WorkerPool/src"
	"time"
)

type SimpleTask struct {

}

func (task SimpleTask) Execute()  {
	time.Sleep(10*time.Second)
	fmt.Println("completed request!")
	fmt.Println("sleeped 3 seconds")
}

func main() {
	workerpool.Start(5)

	http.HandleFunc("/work", work)

	// start worker pool with 5 workers working together


	// Start the HTTP server!
	fmt.Println("HTTP server listening on 127.0.0.1:8000")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}

func work(w http.ResponseWriter, r *http.Request) {

	fmt.Println("started new request");

	newTask := SimpleTask{}

	//newTask.Execute();

	workerpool.AddNewTask(newTask)

	w.WriteHeader(http.StatusOK)
}
