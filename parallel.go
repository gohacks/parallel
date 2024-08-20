package parallel

import (
	"context"
	"sync"
)

type Parallel struct {
	maxGoroutines int
	sem           chan struct{}
	wg            sync.WaitGroup
	ctx           context.Context
	cancel        context.CancelFunc
	errors        []error
	mu            sync.Mutex
}

// New creates a new instance of Pool with the specified limit of goroutines.
func New(ctx context.Context, maxGoroutines int, cancelOnError bool) *Parallel {
	var cancel context.CancelFunc
	if cancelOnError {
		ctx, cancel = context.WithCancel(ctx)
	}
	return &Parallel{
		maxGoroutines: maxGoroutines,
		sem:           make(chan struct{}, maxGoroutines),
		cancel:        cancel,
		errors:        make([]error, 0),
		ctx:           ctx,
	}
}

// Run adds a new function to be executed in the parallel with concurrency control.
func (p *Parallel) Run(task func() error) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.sem <- struct{}{} // Blocks if the goroutine limit is reached

		if p.ctx.Err() != nil { // Verifica se o contexto já está cancelado
			return
		}

		err := task() // Executes the provided function
		if err != nil {
			p.mu.Lock()
			p.errors = append(p.errors, err) // Append errors
			p.mu.Unlock()

			if p.cancel != nil {
				p.cancel() // Cancell all goroutines if cancelOnError for true
			}
		}

		<-p.sem // Releases after execution
	}()
}

// Wait aguarda que todas as goroutines terminem.
func (p *Parallel) Wait() {
	p.wg.Wait()
}

func (p *Parallel) Errors() []error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]error{}, p.errors...)
}
