

// Using a new Keywords in a golang used apointers direacatally used a new keywords
package main

import "fmt"

func main() {
	nums := new(int) 

	*nums = 42
	fmt.Println("The value of nums is:", *nums) // Dereferencing to get the value
}