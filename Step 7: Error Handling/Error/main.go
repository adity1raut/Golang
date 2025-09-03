

package main 

import ( "fmt")


func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    defer fmt.Println("Cleanup: defer called")
	fmt.Println("Executing safeFunction")
    panic("Unexpected error!")
}

func main() {
    safeFunction()
    fmt.Println("Program continues after recovery")
}
