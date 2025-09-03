


package main

import "fmt"

func main() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
	defer fmt.Println("This runs second")
	fmt.Println("This runs third")
	fmt.Println("This runs fourth")
	fmt.Println("This runs fifth")
	defer fmt.Println("This runs sixth")
}


func test() int {
    defer fmt.Println("Deferred runs before return")
    return 100
}

func main() {
    fmt.Println("Returned:", test())
}
