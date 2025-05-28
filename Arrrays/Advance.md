package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    var num int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&num)

    nums := make([]int, num)
    for i := range nums {
        fmt.Scan(&nums[i])
    }

    // Convert int slice to string slice for joining
    strNums := make([]string, num)
    for i, v := range nums {
        strNums[i] = strconv.Itoa(v)
    }
    fmt.Println(strings.Join(strNums, " ")) // Output: 1 2 3
}