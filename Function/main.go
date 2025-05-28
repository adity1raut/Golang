package main 

import "fmt"

func main () {
	// Function to add two integers
	add := func(a int, b int) int {
		return a + b
	}

	// Function to multiply two integers
	multiply := func(a int, b int) int {
		return a * b
	}

	// Using the functions
	sum := add(5, 3)
	product := multiply(4, 6)

	fmt.Println("Sum:", sum)          // Output: Sum: 8
	fmt.Println("Product:", product)  // Output: Product: 24
}