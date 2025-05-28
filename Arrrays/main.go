package main

import "fmt"

// Define a custom array type
type IntArray [5]int

// Method to calculate the sum of elements
func (arr IntArray) Sum() int {
    total := 0
    for _, v := range arr {
        total += v
    }
    return total
}

// Method to find the maximum element
func (arr IntArray) Max() int {
    max := arr[0] // Assume first element is max
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
    // Create an instance of the custom array
    numbers := IntArray{1, 2, 3, 4, 5}

    // Call methods
    fmt.Println("Sum:", numbers.Sum()) // Output: Sum: 15
    fmt.Println("Max:", numbers.Max()) // Output: Max: 5
}