// Gettting A Array in A Function and Passing The Argurments in  and A function through an arrays

package main 

import "fmt"

func isVoted(nums []int) int {

    count := 0 

    for _, age := range nums {
        if age >= 18 {
            count++
        } 
    }
    return count
}

func main() {
    nums := []int {75, 25 , 65, 12 ,11, 45, 78, 90, 100, 50}

    candidate := isVoted(nums)
    fmt.Print("Checking if the candidate is eligible to vote...\n")
    fmt.Println("The candidate is eligible to vote:", candidate)
}