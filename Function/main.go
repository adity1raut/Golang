package main

import (
	"fmt"
)

func main () {
	fmt.Println("Hellow Aditya Welcome To Golang!")

	var num int 
	fmt.Println("Enter Your Age !")
	fmt.Scan(&num)

	result := Voting(num)

	if result {
		fmt.Println("You Can Vote ")
	}else {
		fmt.Println("You Not Eligible For Voting !")
	}


}

func Voting (age int ) bool {
	if age > 18 {
		return true
	}else {
		return false
	}
}