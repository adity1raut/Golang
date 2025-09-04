package main

import (
	"fmt"
	"strings"
)


func main(){

	var s string = "Hello, World!"
	println(s) // Output: Hello, World

	var s2 string = "World"

    ss := strings.Compare(s,s2)
	print(strings.Contains(s , s2))
	if ss == 0 {
		fmt.Println("Strings are equal !")
	} else if ss < 0 {
		fmt.Println("String Are Less Than!")
	}else {
		fmt.Println("String and Greater Than !")
	}

	fmt.Print(strings.Contains(s , s2))
}