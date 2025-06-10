// package main

// import (
//     "fmt"
//     "sync"
// )

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for job := range jobs {
//         fmt.Printf("Worker %d processing job %d\n", id, job)
//     }
// }

// func main() {
//     jobs := make(chan int, 100)
//     var wg sync.WaitGroup

//     // 5 workers start karo (Fan-Out)
//     for i := 1; i <= 5; i++ {
//         wg.Add(1)
//         go worker(i, jobs, &wg)
//     }

//     // Jobs bhejo
//     for j := 1; j <= 20; j++ {
//         jobs <- j
//     }
//     close(jobs)

//     wg.Wait()
// }