package main

import (
	"fmt"
	"strings"
)


func main(){

	var s string = "Hello, World!"
	println(s) // Output: Hello, World!
	var s1 string = "Hello, " + "World!"

	println(s1) // Output: Hello, World!
	println(len(s)) // Output: 13

	// strings.Contains(s, "World") // Check if "World" is in the string
	println(strings.Contains(s, s1)) // Output: true

	

	s2 := strings.Split(s, ", ") // Split the string by ", "
	fmt.Print(s2) // Output: [Hello World!]
}