package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 10)

	// Start 3 worker goroutines (Fan-Out)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}

	// Send 9 jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs) // Close the channel so workers exit
	time.Sleep(4 * time.Second) // Wait for workers to finish
}
