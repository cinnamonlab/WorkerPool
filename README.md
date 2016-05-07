# WorkerPool
The worker pool written in golang

- To start workerPool, just run it with number of workers will be working togerther
```
workerpool.Start(5)
```

- After started workerpool. You can add new task to worker by create new task object which implement Execute() function from ITask interface.

```
type SimpleTask struct {

}

func (task SimpleTask) Execute()  {
	time.Sleep(10*time.Second)
	fmt.Println("completed request!")
	fmt.Println("sleeped 10 seconds")
}
```
```
newTask := SimpleTask{}
workerpool.WorkPool <- newTask
```


