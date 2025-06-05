package main

import "fmt"

func SayHello(ch chan string) {
	ch <- "Hello From a Go Routine"
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	ch := make(chan string)
	go SayHello(ch)
	fmt.Println(<-ch) // Receive and print message from SayHello

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Important: close jobs channel after sending

	// Collect results
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}
