

package main 

import "fmt"

func main () {
    slice := make([]int, 3, 5) // Length=3, Capacity=5
    fmt.Println(slice)         // Output: [0 0 0]
    fmt.Println(len(slice))    // Output: 3
    fmt.Println(cap(slice))    // Output: 5

    c := make([]int , len(slice))
    copy(c , slice)
    
    c = append(c,4 , 5 , 8 , 9 , 7 , 40 ,14 , 45 )
    println(c)
}