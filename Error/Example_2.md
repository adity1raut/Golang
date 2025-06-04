package main

import "fmt"

func riskyFunction() {
    defer func() {
        r := recover()
        if r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("About to panic...")
    panic("Something failed")
    fmt.Println("This will not run")
}

func main() {
    riskyFunction()
    fmt.Println("Program continues after recover")
}
