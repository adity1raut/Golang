

//Getting  a Normals Examples in a Go inbuild Functions using "fmt" package

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