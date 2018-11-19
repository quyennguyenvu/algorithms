package workerpool

import (
	"fmt"
	"strconv"
	"time"
)

type SecondJob struct {
	Task string
}

func (j *SecondJob) Expose() {
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

func (j *SecondJob) fetchData(models chan model) {
	for i := 1; i < 20; i++ {
		var m model
		m = &SecondJob{Task: "second job: task no." + strconv.Itoa(i)}
		models <- m
	}
	close(models)
}

func (j *SecondJob) doSomething() {
	fmt.Println(j.Task)
	time.Sleep(2 * time.Second)
}
