	package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from goroutine")
		time.Sleep(5000 * time.Millisecond)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Main interactions")
		time.Sleep(500 * time.Millisecond)
	}

	go sayHello() // run as a goroutine

	for i := 0; i < 5; i++ {
		fmt.Println("Main function")
		time.Sleep(500 * time.Millisecond)
	}
}
