package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a // p holds the address of a

	fmt.Println(*p) // Output: 10

	*p = 20
	fmt.Println(a) // Origanal variablle ko direacatally update kiya ja sakata hai

	var ptr *int
	fmt.Println(ptr) // Output: <nil>

}
