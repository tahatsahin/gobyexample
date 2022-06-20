package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	// waitgroup will help us wait for all goroutines to finish
	var wg sync.WaitGroup

	// start 50 goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// to atomically increment the counter use AddUint64, giving
		// it the memory address of ops counter.
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				// fmt.Println(atomic.LoadUint64(&ops))
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}
