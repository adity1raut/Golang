package main 

import "fmt"

func main () {
	var num int 
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	facto := Factorial(num)
	fmt.Println(facto)

	sum := Addition(num)
	fmt.Println(sum)
}


func Factorial (a int) int {
	if a == 0 {
		return 1
	}
	return a * Factorial(a-1)
}

func Addition (num int ) int {
	if num == 0 {
		return 0
	}
	return num + Addition(num-1)
}