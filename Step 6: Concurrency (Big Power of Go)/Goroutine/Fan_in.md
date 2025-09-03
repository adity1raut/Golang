package main

import (
	"fmt"
	"time"
)

// Producer 1
func producer1(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Producer 1: message %d", i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Producer 2
func producer2(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Producer 2: message %d", i)
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string)

	// Start two producers (Fan-In)
	go producer1(ch)
	go producer2(ch)

	// Consumer
	for i := 0; i < 10; i++ {
		msg := <-ch
		fmt.Println("Received:", msg)
	}
}
