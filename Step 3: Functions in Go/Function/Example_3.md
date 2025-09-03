``` go

package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// This program demonstrates a closure in Go.
// The `intSeq` function returns another function that increments and returns an integer each time it is called.
// The returned function captures the variable `i` from its surrounding context, allowing it to maintain state across calls.
// The main function calls the returned function multiple times, printing the incremented values.
// The output will show the sequence of integers starting from 1.
// Output:

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
	fmt.Println(nextInt())
    fmt.Println(nextInt())
}