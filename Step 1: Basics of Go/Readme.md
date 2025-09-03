
## What is Go?

Go, often referred to as Golang, is a statically typed, compiled programming language designed at Google. It was created to improve programming productivity in an era of multicore processors, networked machines, and large codebases.

Go is known for its "less is more" philosophy. It has a simple, clean syntax and a small set of features, which makes it easy to learn and read.

## Why Use Go?

Go is popular for building backend services, APIs, command-line tools, and network applications due to three core strengths:

- **Speed 🚀**: Go is a compiled language, which means your code is translated directly into machine code that the processor can execute. This makes it significantly faster than interpreted languages like Python or JavaScript, which are read line-by-line by an interpreter at runtime.

- **Concurrency 🔄**: Concurrency is the ability to have multiple tasks running at the same time. Go has built-in features called goroutines and channels that make it incredibly easy to write concurrent programs. A goroutine is like a very lightweight thread, and you can run thousands of them without much overhead. This makes Go excellent for applications that need to handle many tasks simultaneously, like a web server handling multiple user requests.

- **Simplicity 😊**: The language has a very small specification and only 25 keywords. This simplicity leads to code that is easy to write, read, and maintain. It also comes with a powerful standard library and a robust set of built-in tools for formatting, testing, and managing code.

## Installation and Setup

Installing Go is straightforward.

1. **Download**: Visit the official Go website at [go.dev/dl/](https://go.dev/dl/) and download the installer for your operating system (Windows, macOS, or Linux).

2. **Install**: Run the installer and follow the on-screen instructions. It will set up the Go toolchain and environment variables for you.

3. **Verify**: Open your terminal or command prompt and type `go version`. You should see the installed version of Go printed, like `go version go1.22.5 windows/amd64`.

### `go run` vs. `go build`

Go comes with two primary commands for running your code:

- **`go run`**: This command compiles and runs the specified Go file in one step. It's great for quickly testing a small program. The executable file it creates is temporary and is deleted after the program finishes.

  Example: `go run hello.go`

- **`go build`**: This command only compiles your code. It creates a permanent, executable file (e.g., `hello.exe` on Windows or `hello` on macOS/Linux) in your current directory. You can then run this file anytime without needing to recompile.

  Example: `go build hello.go` followed by `./hello` to run the program.

## Your First Program (Hello, World)

Let's write a "Hello, World!" program. Create a file named `hello.go` and add the following code:

```go
// All Go files start with a package declaration.
// The 'main' package is special; it defines a standalone executable program.
package main

// The 'import' statement is used to include code from other packages.
// The 'fmt' package provides functions for formatted I/O (like printing to the console).
import "fmt"

// The 'main' function is the entry point of every executable program.
// When you run the program, the code inside this function is executed.
func main() {
    fmt.Println("Hello, World")
}
```

To run this, save the file, open your terminal in the same directory, and type:
```
go run hello.go
```

You will see the output: `Hello, World`

## Syntax: Variables, Constants, and Data Types

### Variables

A variable is a named storage location for data. In Go, you can declare variables in a few ways.

- **With `var` keyword**: The standard way is `var <name> <type>`.

```go
// Declares an integer variable named 'age'
var age int
age = 30 // Assign a value

// You can also declare and initialize in one line
var name string = "Alice"
```

- **Short Declaration Operator `:=`**: This is the most common way to declare and initialize a variable inside a function. Go automatically infers the data type from the value.

```go
// This is equivalent to 'var score int = 95'
score := 95

// This is equivalent to 'var isReady bool = true'
isReady := true
```

### Constants

A constant is a variable whose value cannot be changed after it's assigned. They are declared using the `const` keyword.

```go
const Pi float64 = 3.14159
const IsProduction bool = true
const SiteName string = "My Website"

// You cannot change a constant's value
// Pi = 3.14 // This would cause a compilation error
```

### Basic Data Types

Go has several built-in data types.

| Type     | Description                            | Example             |
|----------|----------------------------------------|---------------------|
| `int`    | Whole numbers (integers).              | `10`, `-5`, `1000`  |
| `float64`| Floating-point (decimal) numbers.      | `3.14`, `-0.5`, `99.99` |
| `bool`   | A boolean value, either true or false. | `true`, `false`     |
| `string` | A sequence of characters.              | `"hello"`, `"Go Lang"` |

## Operators

Operators are symbols that perform operations on variables and values.

### Arithmetic Operators

These are used for performing mathematical calculations.

| Operator | Name        | Description                      | Example (a := 10, b := 3) | Result |
|----------|-------------|----------------------------------|---------------------------|--------|
| `+`      | Addition    | Adds two operands                | `a + b`                   | `13`   |
| `-`      | Subtraction | Subtracts second operand         | `a - b`                   | `7`    |
| `*`      | Multiplication | Multiplies both operands      | `a * b`                   | `30`   |
| `/`      | Division    | Divides first by second          | `a / b`                   | `3`    |
| `%`      | Modulus     | Remainder of a division          | `a % b`                   | `1`    |

Example code:

```go
package main

import "fmt"

func main() {
    x := 20
    y := 5

    fmt.Println("x + y =", x + y) // Output: x + y = 25
    fmt.Println("x - y =", x - y) // Output: x - y = 15
    fmt.Println("x * y =", x * y) // Output: x * y = 100
    fmt.Println("x / y =", x / y) // Output: x / y = 4
    fmt.Println("x % y =", 21 % 5) // Output: x % y = 1
}
```

### Comparison Operators

These operators compare two values and return a boolean result (`true` or `false`).

| Operator | Name                     | Example (a := 10, b := 20) | Result  |
|----------|--------------------------|----------------------------|---------|
| `==`     | Equal to                 | `a == b`                   | `false` |
| `!=`     | Not equal to             | `a != b`                   | `true`  |
| `>`      | Greater than             | `a > b`                    | `false` |
| `<`      | Less than                | `a < b`                    | `true`  |
| `>=`     | Greater than or equal to | `a >= 10`                  | `true`  |
| `<=`     | Less than or equal to    | `b <= 15`                  | `false` |

Example code:

```go
package main

import "fmt"

func main() {
    price := 100
    budget := 150

    fmt.Println("Price equals budget?", price == budget)       // false
    fmt.Println("Can afford?", price <= budget)                // true
}
```

### Logical Operators

These operators are used to combine conditional statements.

| Operator | Name        | Description                                      | Example (a := true, b := false) | Result  |
|----------|-------------|--------------------------------------------------|---------------------------------|---------|
| `&&`     | Logical AND | Returns true if both statements are true.        | `a && b`                        | `false` |
| `\|\|`   | Logical OR  | Returns true if at least one statement is true.  | `a \|\| b`                      | `true`  |
| `!`      | Logical NOT | Reverses the result; returns false if true.      | `!b`                            | `true`  |

Example code:

```go
package main

import "fmt"

func main() {
    age := 25
    hasLicense := true

    canDrive := age >= 18 && hasLicense
    fmt.Println("Can this person drive?", canDrive) // true

    isTeenager := age >= 13 && age <= 19
    fmt.Println("Is this person a teenager?", isTeenager) // false

    isNotTeenager := !isTeenager
    fmt.Println("Is this person not a teenager?", isNotTeenager) // true
}
```