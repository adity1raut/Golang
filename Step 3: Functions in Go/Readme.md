
## 📌 Defining and Calling Functions
You define a function using the `func` keyword, followed by:
- Function name  
- Parameters (with types)  
- Return type(s)  
- Function body inside `{}`  

**Syntax:**
```go
func functionName(parameterName type) returnType {
    // function body
}
````

To execute the function, you **call** it by writing its name followed by parentheses.

### Example:

```go
package main

import "fmt"

// Defining the greet function
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Alice") // Output: Hello, Alice
    greet("Bob")   // Output: Hello, Bob
}
```

---

## 📌 Parameters and Return Values

* **Parameters** → Inputs to a function.
* **Return Values** → Outputs of a function.

Go is **pass-by-value** (a copy of the value is passed).

### Example with Return:

```go
package main

import "fmt"

// a and b are parameters, return type is int
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(5, 3)
    fmt.Println("Sum:", sum) // Output: Sum: 8
}
```

---

## 📌 Named Return Values

You can declare return variables in the signature. A `return` without arguments will return their current values.

```go
func subtract(a int, b int) (result int) {
    result = a - b
    return // returns 'result'
}
```

---

## 📌 Multiple Return Values

Functions in Go can return multiple values (commonly a result and an error).

```go
package main

import (
    "errors"
    "fmt"
)

// Returns a float64 and an error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: 5
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: cannot divide by zero
    }
}
```

➡ If you only need one return value:

```go
result, _ := divide(10, 2) // ignore error
```

---

## 📌 Variadic Functions

A variadic function accepts a variable number of arguments (`...` syntax).

```go
package main

import "fmt"

// 'nums' is a slice of ints
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2))          // 3
    fmt.Println(sum(1, 2, 3, 4))    // 10
    numbers := []int{5, 6, 7}
    fmt.Println(sum(numbers...))    // 18
}
```

---

## 📌 Anonymous Functions

Anonymous (unnamed) functions can be:

* Assigned to variables
* Passed as arguments
* Executed immediately

```go
package main

import "fmt"

func main() {
    add := func(a, b int) int {
        return a + b
    }

    result := add(3, 4)
    fmt.Println("Result:", result) // 7
}
```

---

## 📌 Closures

Closures are anonymous functions that capture variables from their surrounding scope.

```go
package main

import "fmt"

func makeIncrementor() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    incrementor := makeIncrementor()

    fmt.Println(incrementor()) // 1
    fmt.Println(incrementor()) // 2
    fmt.Println(incrementor()) // 3

    another := makeIncrementor()
    fmt.Println(another()) // 1 (independent)
}
```

---
