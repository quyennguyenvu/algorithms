package playground

import "algorithms/workerpool"

func GotoWork() {
	provider := workerpool.New()
	workerpool.Expose(provider)
}
