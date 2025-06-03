package main

import "fmt"

func Swap(num1 *int , num2 *int) {
	temp := *num1 
	*num1 = *num2
	*num2 = temp
}

func main () {

	fmt.Println("Enters a Two numbers For swapping ....")
	var a , b int
	fmt.Print("Enter a First number :")
	fmt.Scan(&a)

	fmt.Print("Enter a Secound number :")
	fmt.Scan(&b)

	Swap(&a , &b)

    fmt.Print("Printing a Swapping Numbers :")
	println(a)
	println(b)


}

