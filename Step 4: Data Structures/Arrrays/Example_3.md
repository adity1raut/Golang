# Array Basics in Go

This program demonstrates different ways of working with arrays in Go, including initialization, accessing, modifying, and iterating.

## Code

```go
package main

import "fmt"

func main() {
    // Zero value array (default values are 0)
    var arr1 [5]int
    fmt.Println(arr1)

    // Initialized empty array (explicit zeros)
    arr2 := [5]int{0, 0, 0, 0, 0}
    fmt.Println(arr2)

    // Literal initialization
    arr3 := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr3)

    // Accessing elements
    fmt.Println("Element at index 2:", arr3[2])

    // Modifying an element
    arr3[2] = 10
    fmt.Println("Modified array:", arr3)

    // Length of the array
    fmt.Println("Length of the array:", len(arr3))

    // Iterating over the array
    for i, v := range arr3 {
        fmt.Printf("Index %d: Value %d\n", i, v)
    }

    // Note:
    // In iteration using range:
    // i → represents the index
    // v → represents the value at that index
}
