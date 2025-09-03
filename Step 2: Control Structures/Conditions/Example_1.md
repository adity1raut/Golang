# Swap Two Numbers in Go

This program demonstrates how to swap two numbers in **Golang** using pointers.

---

## 📌 Code

```go
package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Enter a First Number : ")
    fmt.Scan(&num1)

    fmt.Print("Enter a Second Number : ")
    fmt.Scan(&num2)

    Swap(&num1, &num2)

    println("Number 1 :", num1)
    println("Number 2 :", num2)
}

func Swap(num1 *int, num2 *int) {
    temp := *num1
    *num1 = *num2
    *num2 = temp
}
