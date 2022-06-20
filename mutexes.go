package main

import (
	"fmt"
	"sync"
)

// container holds a map of counters, to update it concurrently from
// multiple goroutines we add a mutex to synchronize acces.
// note that mutexes must not be copied, so if this struct is passed
// around, it should be done by a pointer!!
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// lock the mutex before accesing counters; unlock it at the end of the
// function using a defer
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// zero value of a mutex is usable as-is, so no initialization is required
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	// run several goroutines concurrently, 2 of them access same counter
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
