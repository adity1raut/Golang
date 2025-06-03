package main

import "fmt"

func main () {
    // Zero value array (nil, cannot be used directly)
    var arr1 [5]int
    fmt.Println(arr1)

    // Initialized empty array (ready to use)
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

    // In The Topic of itterration in an Arrays the i : is define the inex of and Arrays and v : is define the value of an Arrays
    
}