

package main 

import "fmt"

func evnOdd(num int) {
	if num%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}

func isPrime(num int) {
	if num <= 1 {
		fmt.Println("The number is not prime.")
		return
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			fmt.Println("The number is not prime.")
			return
		}
	}
	fmt.Println("The number is prime.")
}

func isPositive(num int) {
	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}

func isLeapYear(year int) {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("The year is a leap year.")
	} else {
		fmt.Println("The year is not a leap year.")
	}
}
func isPalindrome(s string) {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			fmt.Println("The string is not a palindrome.")
			return
		}
		left++
		right--
	}
	fmt.Println("The string is a palindrome.")
}



func main(){
	fmt.Println("Checking if a number is even or odd:")
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	evnOdd(num)

	fmt.Println("\nChecking if a number is prime:")
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	isPrime(num)

	fmt.Println("\nChecking if a number is positive, negative, or zero:")
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	isPositive(num)

	fmt.Println("\nChecking if a year is a leap year:")
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)
	isLeapYear(year)

	fmt.Println("\nChecking if a string is a palindrome:")
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)
	isPalindrome(str)
	fmt.Println("\nAll checks completed successfully!")
	fmt.Println("Thank you for using the program!")
	fmt.Println("Have a great day!")


}