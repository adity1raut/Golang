package main 

import "fmt"

// Example of Higher-Order Functions in Go
// Higher-order functions are functions that can take other functions as arguments or return them as results.
// In this example, we define a few basic arithmetic operations and a function that applies these operations.
// This demonstrates how functions can be passed around as first-class citizens in Go.
// Example of Higher-Order Functions in Go
// Higher-order functions are functions that can take other functions as arguments or return them as results.
// In this example, we define a few basic arithmetic operations and a function that applies these operations.
// This demonstrates how functions can be passed around as first-class citizens in Go.
// This example demonstrates how to use higher-order functions in Go by defining basic arithmetic operations
// and a function that applies these operations. It shows how functions can be passed as arguments,	
// allowing for flexible and reusable code.
// This example demonstrates how to use higher-order functions in Go by defining basic arithmetic operations


func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	if b == 0 {
		return 0 // Handle division by zero
	}
	return int(a) / int(b)
}

func ApplyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main(){


	r:= ApplyOperation(10, 5, Add)
	fmt.Println("Addition:", r)


	r = ApplyOperation(10, 5, Subtract)
	fmt.Println("Subtraction:", r)

	r = ApplyOperation(10, 5, Multiply)
	fmt.Println("Multiplication:", r)

	r = ApplyOperation(10, 5, Divide)
	fmt.Println("Division:", r)

	r = ApplyOperation(10, 0, Divide)
	fmt.Println("Division by zero:", r)

}