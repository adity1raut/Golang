
```go
package main 

func add(n int) int {
	if n == 0 {
		return 0
	}
	return n + add(n-1)
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	// Demonstrating recursion problems in Go

	r := add(8)
	println("Sum of first 8 numbers is:", r)

	f := Factorial(5)
	println("Factorial of 5 is:", f)

	fib := Fibonacci(6)
	println("Fibonacci of 6 is:", fib)
}
