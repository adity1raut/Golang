package main
import "fmt"

func isOdd(n int) bool {
	return n%2 != 0
}

func Factorial(n int ) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
func main() {
	num := 42

	if isOdd(num) {
		println(num, "is odd")
	} else {
		println(num, "is even")
	}

	print("Enter a number: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		println("Invalid input")
		return
	}

	result := Factorial(input)
	println("Factorial of", input, "is", result)
}