package main

import "fmt"

func searchElement(arr []int, element int) bool {
	for _, value := range arr {
		if value == element {
			return true
		}
	}
	return false
}

func printArray(arr []int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

// main function to demonstrate array operations

func main() {

	fmt.Print("Enter the size of the array: ")
	var size int
	fmt.Scan(&size)

	arr := make([]int, size)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Search The Elemetn in an Arrays :")
	var element int
	fmt.Scan(&element)
	if searchElement(arr, element) {
		fmt.Println("Element found in the array.")
	} else {
		fmt.Println("Element not found in the array.")
	}

    fmt.Println("Elements of the array:")
	printArray(arr)
}
