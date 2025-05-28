package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&num)

    // Create a slice with size 'num'
    nums := make([]int, num)

    // Input elements into the slice
    for i := 0; i < num; i++ {
        fmt.Scan(&nums[i])
    }

    // Print elements
    for i := 0; i < num; i++ {
        if i == num-1 {
            fmt.Print(nums[i])
        } else {
            fmt.Print(nums[i], " ")
        }
    }



	// Print a newline for better readability
	fmt.Println() // Print a newline after the elements
	fmt.Println(nums) // Print a newline at the end
}