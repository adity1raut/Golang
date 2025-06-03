package main

import "fmt"

func main() {

	var num1 , num2 int = 10, 20

	fmt.Println("Before swapping: num1 =", num1, ", num2 =", num2)
	swap(&num1, &num2)

	fmt.Println("After swapping: num1 =", num1, ", num2 =", num2)
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}