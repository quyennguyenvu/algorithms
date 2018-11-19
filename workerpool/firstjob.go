package workerpool

import (
	"fmt"
	"strconv"
	"time"
)

type FirstJob struct {
	Task string
}

func (j *FirstJob) Expose() {
	fmt.Println("Exposing first job ...")
	startTime := time.Now()
	noWorkers := 10

	var models = make(chan model, 100)
	go j.fetchData(models)
	Pool(noWorkers, models)

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func (j *FirstJob) fetchData(models chan model) {
	for i := 1; i < 20; i++ {
		var m model
		m = &FirstJob{Task: "first job: task no." + strconv.Itoa(i)}
		models <- m
	}
	close(models)
}

func (j *FirstJob) doSomething() {
	fmt.Println(j.Task)
	time.Sleep(2 * time.Second)
}
