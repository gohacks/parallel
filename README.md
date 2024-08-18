# parallel

The parallel library provides a simple way to manage concurrent execution of tasks using a pool of goroutines with a specified limit. This allows you to control and limit the number of concurrent goroutines in your Go applications.

## Features
- Create a pool of goroutines with a maximum concurrency limit.
- Add tasks to the pool with concurrency control.
- Wait for all tasks to complete.

## Installation
To use the parallel library in your Go project, simply import parallel package into your project.

```
go get github.com/gohacks/parallel
```

## Usage

### Example
Here's a basic example of how to use the parallel library:

```
package main

import (
	"fmt"
	"time"
	"github.com/gohacks/parallel"
)

func main() {
	// Create a new pool with a maximum of 3 concurrent goroutines
	pool := parallel.New(3)

	// Define a task function
	func() {
		
	}

	// Add tasks to the pool
	for i := 1; i <= 10; i++ {
		pool.Run(func() {
            yourFunction(yourParams) // call your functions here
	    })
	}

	// Wait for all tasks to complete
	pool.Wait()
}
```

## API

New(maxGoroutines int) *Parallel
Creates a new instance of Parallel with the specified limit of goroutines.

maxGoroutines: The maximum number of goroutines that can run concurrently.
(*Parallel) Run(task func())
Adds a new task to be executed in the parallel with concurrency control.

task: A function to be executed by the parallel.
(*Parallel) Wait()
Waits for all tasks to complete.

## Notes
- The library uses a semaphore pattern with channels to manage concurrency.
- Ensure that the maxGoroutines value is set according to the resources available and the requirements of your application.

## Contributing
Contributions are welcome! If you have any improvements or bug fixes, please open an issue or submit a pull request.

## License
This library is released under the MIT License. See the [LICENSE](./LICENSE) file for more details.