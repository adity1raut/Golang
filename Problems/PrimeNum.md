

package main

import "fmt"
// This program checks if a number is prime or not

func main() {
	var num int = 7
	var isPrime bool = true

	if num < 2 {
		isPrime = false
	} else {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println("Number is Prime")
	} else {
		fmt.Println("Number is Not Prime")
	}
}