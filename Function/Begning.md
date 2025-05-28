package main 

import "fmt"


func main () {

	num := add(10 ,20)
	fmt.Println("Sum of 10 and 20 is : ", num)

	fact := Factorial(64)
	fmt.Println("Factorial of 5 is : ", fact)

	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient is : ", quotient)
	fmt.Println("Remainder is : ", remainder)
}


func Factorial (n int) int {
	if n == 0 ||  n == 1 {
		// Factorial of 0 and 1 is 1	
		return 1
	}
	return n * Factorial(n-1)
}

//Bhai return karnese pahale hame wahape type bhi mension karna padta hai jaise ki baki language me karte hai isliy to used karte hai functions 


func add (a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
