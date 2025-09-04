# Go Input Example (Reading a Full Line)

This example demonstrates how to read a **full line of input** in Go using the `bufio` package.  
It allows users to enter a sentence (including spaces), and captures it as a string.

## Example Code

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Enter a line of text (press Enter to finish):")

    // Create a new scanner to read input from standard input
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Input: ")

    // Read the input line
    if scanner.Scan() {
        input := scanner.Text()
        fmt.Println("You entered:", input)
    }
}
