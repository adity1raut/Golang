package main 

import "fmt"

func main() {

	nums := [5]int{10, 20, 30, 40, 50}
	Increment(&nums)

	for _ , num := range nums {
		fmt.Print(num , " ")
	}

}

func Increment(nums *[5]int) {
	for i := 0; i < len(nums); i++ {
		nums[i] += 10
	}
}
