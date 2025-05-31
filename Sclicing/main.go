package main

import "fmt"

func main() {

	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)      // Output: [1 2 3 4 5 6 7 8 9]
	fmt.Println(len(nums)) // Output: 9
	fmt.Println(cap(nums)) // Output: 9

	slice := make([]int, 3, 5) // Length=3, Capacity=5
	fmt.Println(slice)         // Output: [0 0 0]
	fmt.Println(len(slice))    // Output: 3
	fmt.Println(cap(slice))    // Output: 5
	slice[0] = 49
	slice[1] = 56
	slice[2] = 789
	fmt.Println(slice) // Output: [49 56 789]

	c := make([]int, len(nums))
	copy(c, nums)  // Direactally copy karega re bhai tu kuch bhi karle bhai hoga apne app

	fmt.Println(c)
	c = append(c, slice...) // Appending the slice to c (Yanike Add Honge re Bhai Size tu kucha bhi declare kar le vo extra ke hi add honge Mera yasu yasu!)
	fmt.Println(c)

}
