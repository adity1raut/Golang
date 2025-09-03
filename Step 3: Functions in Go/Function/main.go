

package main 

import "fmt"

func isVoted(age int) {
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}

func main( )  {

	fmt.Println("Chacking a user Can get Vote or Not !")
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	
	isVoted(age)
}