package main

import (
	"fmt"
	"os"
)


func main () {
	file , err := os.Open("main.txt")

	if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() 

	fmt.Println("File opened successfully:", os.WriteFile("main.txt", []byte("Hello, World!"), 0644))

	fileInfo := make([]byte, 100) // Create a buffer of 100 bytes
    count, err := file.Read(fileInfo)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Step 4: Print file content
    fmt.Printf("File Content (%d bytes):\n%s\n", count, fileInfo[:count])
}