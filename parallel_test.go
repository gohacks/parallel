package parallel

import (
	"sync"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	const maxGoroutines = 5
	const numTasks = 10

	pool := New(maxGoroutines)

	var mu sync.Mutex
	taskCount := 0

	// This function will increment taskCount and simulate a long-running task
	task := func() {
		time.Sleep(100 * time.Millisecond) // Simulate job
		mu.Lock()
		taskCount++
		mu.Unlock()
	}

	for i := 0; i < numTasks; i++ {
		pool.Run(task)
	}

	pool.Wait()

	if taskCount != numTasks {
		t.Errorf("The completed task count is %d, but we expected %d", taskCount, numTasks)
	}
}

func TestParallelConcurrencyLimit(t *testing.T) {
	const maxGoroutines = 5
	const numTasks = 50

	pool := New(maxGoroutines)

	var mu sync.Mutex
	currentlyRunning := 0
	maxRunning := 0

	task := func() {
		mu.Lock()
		currentlyRunning++
		if currentlyRunning > maxRunning {
			maxRunning = currentlyRunning
		}
		mu.Unlock()

		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		currentlyRunning--
		mu.Unlock()
	}

	for i := 0; i < numTasks; i++ {
		pool.Run(task)
	}

	pool.Wait()

	if maxRunning > maxGoroutines {
		t.Errorf("The maximum number of concurrently running goroutines was %d, but it should not exceed %d", maxRunning, maxGoroutines)
	}
}
