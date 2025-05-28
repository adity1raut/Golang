package main

import (
	"fmt"
)

func main() {
	var nums int = 45
	var num string = "Aditya Waman Raut"
	if nums >= 18 {
		fmt.Println("User Can Vote !")
		fmt.Println(num)
	} else {
		fmt.Println("User Don't Vote !")
		fmt.Println(num)
	}
}