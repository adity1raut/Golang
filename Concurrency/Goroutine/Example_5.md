package main

import (
	"fmt"
	"sync"
	"time"
)



func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "started")
	time.Sleep(time.Second)
	fmt.Println("Task", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed")
}
