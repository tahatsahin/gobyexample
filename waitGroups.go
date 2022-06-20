package main

import (
	"fmt"
	"sync"
	"time"
)

// worker instance
func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

// this waitgroup is used to wait for all the goroutines launched here
// to finish. if a waitGroup is explicitly passed into functions, it
// should be done by pointer.
func main() {
	var wg sync.WaitGroup

	// launch several goroutines, increment wg counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// avoid re-use of the same i value in each goroutine closure.
		i := i

		// wrap the worker call in a closure that makes sure to tell
		// the waitGroup that this worker is done. This way the worker
		// does not have to be aware of the concurrency primitives
		// involved in its execution
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// block until wg counter goes back to 0.
	wg.Wait()

	// for more advanced use cases, consider using the errgroup package.
	// https://pkg.go.dev/golang.org/x/sync/errgroup
}
