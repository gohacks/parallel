package parallel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	const maxGoroutines = 5
	const numTasks = 10

	parallel := New(context.Background(), maxGoroutines, false)

	var mu sync.Mutex
	taskCount := 0

	// This function will increment taskCount and simulate a long-running task
	task := func() error {
		time.Sleep(100 * time.Millisecond) // Simulate job
		mu.Lock()
		taskCount++
		mu.Unlock()
		return nil
	}

	for i := 0; i < numTasks; i++ {
		parallel.Run(task)
	}

	parallel.Wait()

	if taskCount != numTasks {
		t.Errorf("The completed task count is %d, but we expected %d", taskCount, numTasks)
	}
}

func TestParallelConcurrencyLimit(t *testing.T) {
	const maxGoroutines = 5
	const numTasks = 50

	parallel := New(context.Background(), maxGoroutines, false)

	var mu sync.Mutex
	currentlyRunning := 0
	maxRunning := 0

	task := func() error {
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
		return nil
	}

	for i := 0; i < numTasks; i++ {
		parallel.Run(task)
	}

	parallel.Wait()

	if maxRunning > maxGoroutines {
		t.Errorf("The maximum number of concurrently running goroutines was %d, but it should not exceed %d", maxRunning, maxGoroutines)
	}
}

func TestParallelWithContextCancelled(t *testing.T) {
	const maxGoroutines = 5
	const numTasks = 10

	pool := New(context.Background(), maxGoroutines, true)

	var mu sync.Mutex
	taskCount := 0

	// This function will increment taskCount and simulate a long-running task
	task := func() error {
		time.Sleep(100 * time.Millisecond) // Simulate job
		if taskCount > 3 {
			return fmt.Errorf("error test")
		}
		mu.Lock()
		taskCount++
		mu.Unlock()
		return nil
	}

	for i := 0; i < numTasks; i++ {
		pool.Run(task)
	}

	pool.Wait()

	if taskCount == numTasks {
		t.Errorf("The completed task count is %d, but we expected not completed tasks", taskCount)
	}
}
