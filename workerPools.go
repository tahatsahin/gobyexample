package main

import (
	"fmt"
	"time"
)

// worker instance. receive work on the jobs channel, send the
// corresponding result on results.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// Worker Pool implementation using go routines and channels
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 3 worker, blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Sending 5 jobs and then closing that channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collecting results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
