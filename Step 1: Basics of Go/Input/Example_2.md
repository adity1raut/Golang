

//Gettting  a Inputs as a full lin e of String with a backspace

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
    // Use the scanner to read a full line of input

    // Read the input line
    if scanner.Scan() {
        input := scanner.Text()
        fmt.Println("You entered:", input)
    }

    
}