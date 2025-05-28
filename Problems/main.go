package main

import "fmt"


func main () {
	// Finding a Greater Number Amoung Three Numbers 

	var num1 , num2 , num3 int 

	fmt.Print("Enter a First Number :")
	fmt.Scan(&num1)

	fmt.Print("Enter a Secound Number :")
	fmt.Scan(&num2)

	fmt.Print("Enter a Third Number :")
	fmt.Scan(&num3)

	Greater := GreaterNum(num1 , num2 , num3)

	fmt.Println(Greater)

}

func GreaterNum ( a int , b int , c int) int {
	num := 0
	if a > b && a > c {
		num = a
		fmt.Println(a , "IS a Greater Number")
	}else if   b > a && b > c {
		num =b 
		fmt.Println(b , "IS a Greater Number")
	} else {
		num = c 
		fmt.Println(c , "IS a Greater Number")
	}
	return num
}