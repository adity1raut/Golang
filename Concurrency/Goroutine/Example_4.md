package main

import (
	"fmt"
	"time"
)

// this is a parrrallerl tasks and this is a h tasks run at the same time. The program waits using Sleep (but you should use sync.WaitGroup instead).
func task(name string) {
	fmt.Println(name, "started")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "done")
}

func main() {
	go task("A")
	go task("B")
	time.Sleep(2 * time.Second)
}
