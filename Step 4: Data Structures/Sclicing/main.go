package main

import "fmt"

func main () {
   nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

   slice := make([]int,len(nums)) // Create a slice with a capacity of 5

   copy(slice, nums[0:]) // Copy the first 5 elements from nums to slice
   fmt.Println("First Slicing " ,slice)

   slice = append(slice, 11, 12, 13) // Append elements to the slice
   fmt.Println("Secound Slicing" ,slice)

   slice = append(slice, 14, 15) // Append more elements to the slice
   fmt.Println("Third Slicing" ,slice)

   slice = append(slice, 16, 17, 18, 19, 20) // Append even more elements
   fmt.Println("Final Slicing" ,slice)
  
}