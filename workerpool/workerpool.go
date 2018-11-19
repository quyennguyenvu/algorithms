package workerpool

import "sync"

// Provider interface
type Provider interface {
	Expose()
}

type model interface {
	doSomething()
}

func Worker(wg *sync.WaitGroup, models chan model) {
	for m := range models {
		m.doSomething()
	}
	wg.Done()
}

func Pool(noWorkers int, models chan model) {
	var wg sync.WaitGroup
	for i := 0; i < noWorkers; i++ {
		wg.Add(1)
		go Worker(&wg, models)
	}
	wg.Wait()
}

// New func
func New() []Provider {
	var first *FirstJob
	var second *SecondJob
	var providers = []Provider{
		first,
		second,
	}
	return providers
}

// Save func
func Expose(providers []Provider) {
	for _, p := range providers {
		p.Expose()
	}
}
