package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when finished

	fmt.Printf("Worker %d: started\n", id)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Worker %d: finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add a goroutine to wait for
		go worker(i, &wg)
	}

	wg.Wait() // Block until all workers finish
	fmt.Println("All workers done.")
}
