package main

import "fmt"

// Function to increment an integer via pointer
func increment(n *int) {
    *n = *n + 1
}

// Define a Person struct
type Person struct {
    Name string
    Age  int
}

// Method with value receiver
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

// Method with pointer receiver
func (p *Person) Birthday() {
    p.Age++
}

// Define Address struct for embedding example
type Address struct {
    City string
}

// Define Employee struct with embedding
type Employee struct {
    Person
    Address
    EmployeeID int
}

func main() {
    // Example 1: Pointers with basic types
    x := 10
    fmt.Println("Original x:", x) // Output: 10

    p := &x
    fmt.Println("Pointer p:", p) // Output: address of x
    fmt.Println("Value at p:", *p) // Output: 10

    *p = 20
    fmt.Println("Modified x:", x) // Output: 20

    increment(p)
    fmt.Println("After increment:", x) // Output: 21

    // Example 2: Structs and pointers
    person := Person{Name: "Alice", Age: 30}
    fmt.Println("Original person:", person) // Output: {Alice 30}

    ptr := &person
    fmt.Println("Pointer to person:", ptr) // Output: &{Alice 30}

    fmt.Println("Name via pointer:", ptr.Name) // Output: Alice
    fmt.Println("Age via pointer:", ptr.Age)   // Output: 30

    ptr.Age = 31
    fmt.Println("After modifying age:", person) // Output: {Alice 31}

    // Example 3: Methods on structs
    person.Greet() // Output: Hello, my name is Alice

    ptr.Birthday()
    fmt.Println("After birthday:", person) // Output: {Alice 32}

    // Example 4: Slice of structs
    people := []Person{
        {Name: "Bob", Age: 25},
        {Name: "Charlie", Age: 28},
    }

    for _, p := range people {
        p.Greet()
    }
    // Output:
    // Hello, my name is Bob
    // Hello, my name is Charlie

    // Modify a struct in the slice via pointer
    ptrToBob := &people[0]
    ptrToBob.Birthday()
    fmt.Println("After Bob's birthday:", people[0]) // Output: {Bob 26}

    // Example 5: Struct embedding
    e := Employee{
        Person:     Person{Name: "David", Age: 40},
        Address:    Address{City: "New York"},
        EmployeeID: 123,
    }

    fmt.Println("Employee:", e.Name, e.City, e.EmployeeID) // Output: David New York 123

    // Can also call methods from embedded struct
    e.Greet() // Output: Hello, my name is David
}