

package main

import "fmt"
// This program checks if a number is prime or not

func main() {
	var num int
	fmt.Print("Enter a number We Whant Sum Of This N numbers : ")	
	fmt.Scan(&num)
	var sum int = 0

	for i := 1; i <= num ; i++ {
		sum += i
	}
	
	fmt.Println("Sum of first", num, "numbers is:", sum)

}