package main

import (
    "fmt"
    "sync"
)

func worker(id int, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 3; i++ {
        results <- fmt.Sprintf("Result from worker %d: %d", id, i)
    }
}

func main() {
    results := make(chan string, 100)
    var wg sync.WaitGroup

    // 3 workers start karo
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, results, &wg)
    }

    // Fan-In: Results ko ek channel mein collect karo
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect kiye hue results ko print karo
    for result := range results {
        fmt.Println(result)
    }
}