

package main 

import "fmt"

func main() {

	num  := 10 
	// assign a Value Using a Pointers in a golang

	var p *int = &num 

	*p = 15 

	fmt.Println("Address" ,p)
	fmt.Println("Values :" , num )
	
}