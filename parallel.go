package parallel

import (
	"sync"
)

type Parallel struct {
	maxGoroutines int
	sem           chan struct{}
	wg            sync.WaitGroup
}

// New creates a new instance of Pool with the specified limit of goroutines.
func New(maxGoroutines int) *Parallel {
	return &Parallel{
		maxGoroutines: maxGoroutines,
		sem:           make(chan struct{}, maxGoroutines),
	}
}

// Run adds a new function to be executed in the parallel with concurrency control.
func (p *Parallel) Run(task func()) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.sem <- struct{}{} // Blocks if the goroutine limit is reached

		task() // Executes the provided function

		<-p.sem // Releases after execution
	}()
}

// Wait aguarda que todas as goroutines terminem.
func (p *Parallel) Wait() {
	p.wg.Wait()
}
