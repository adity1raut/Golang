
package main

import "fmt"

func main () {
    // Finding Greatest Number of and Array in go lang

    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    max := arr[0] // Assume the first element is the largest
    for _, num := range arr {
        if num > max {
            max = num // Update max if a larger number is found
        }
    }
    fmt.Println("The largest number in the array is:", max)
    // Finding Smallest Number of and Array in go lang
    min := arr[0] // Assume the first element is the smallest
    for _, num := range arr {
        if num < min {
            min = num // Update min if a smaller number is found
        }
    }
    fmt.Println("The smallest number in the array is:", min)
}   

// for _, num := range arr
// range arr goes through each element in the slice arr.
// _ ignores the index of the element (we don’t need it here).
// num is the current element of the array in each iteration.