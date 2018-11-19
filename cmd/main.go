package main

import "algorithms/workerpool"

func main() {
	provider := workerpool.New()
	workerpool.Expose(provider)
}
