# Go Built-in Printing Functions (fmt Package)

This example demonstrates how to use Go's built-in functions from the **`fmt`** package to print values in different ways.

## Example Code

```go
package main

import "fmt"

func main() {
    // Printing a string
    fmt.Println("Hello, World!")

    // Printing a formatted string
    name := "Alice"
    age := 30
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)

    // Printing multiple values
    a, b := 5, 10
    fmt.Println("The values are:", a, b)

    // Printing with formatting options
    fmt.Printf("Value of a: %d, Value of b: %d\n", a, b)

    // Printing a floating-point number
    f := 3.14159
    fmt.Printf("The value of pi is approximately %.2f\n", f)
}
