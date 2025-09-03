
package main

import (
    "errors"
    "fmt"
)

var ErrPermission = errors.New("permission denied")

func doSomething() error {
    return fmt.Errorf("operation failed: %w", ErrPermission)
}

func main() {
    err := doSomething()
    if errors.Is(err, ErrPermission) {
        fmt.Println("Access error")
    }
}
