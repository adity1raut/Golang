
## Control Flow Statements

Control flow statements determine the order in which code is executed. Go provides several statements to manage conditions, loops, and jumps.

### Conditional Statements

Conditional statements execute different blocks of code based on whether a condition is `true` or `false`.

#### `if`, `else if`, `else`

This is the most common way to handle conditions.

- `if`: Executes a block of code if its condition is `true`.
- `else if`: If the preceding `if` condition is `false`, this next condition is checked. You can have multiple `else if` blocks.
- `else`: If all preceding `if` and `else if` conditions are `false`, this block is executed.

**Syntax Rules:**
- Parentheses `()` around the condition are not used.
- Curly braces `{}` are **always** required, even for a single-line statement.

```go
package main

import "fmt"

func main() {
    score := 78

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C") // This block will be executed
    } else {
        fmt.Println("Grade: F")
    }
}
```

**Output:**
`Grade: C`

A unique feature in Go is that `if` statements can have a short **initialization statement** before the condition. Variables declared here are only available within the scope of the `if-else` block.

```go
// 'num' is declared and is only accessible inside this if/else block
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
```

#### `switch`

The `switch` statement is a cleaner way to write a long sequence of `if-else if` statements. It compares an expression against a series of `case` values.

**Key Features in Go:**
- **No Fallthrough**: Unlike languages like C or Java, a `case` in Go automatically `break`s. Code does not "fall through" to the next case.
- **Multiple Expressions**: You can list multiple expressions in a single `case`, separated by commas.
- **Expressionless Switch**: A `switch` without an expression is a clean way to write `if-else` logic.

```go
package main

import "fmt"
import "time"

func main() {
    day := time.Now().Weekday()

    switch day {
    case time.Saturday, time.Sunday: // Case with multiple values
        fmt.Println("It's the weekend! 🎉")
    case time.Monday:
        fmt.Println("Manic Monday!")
    default: // This runs if no other case matches
        fmt.Println("It's a weekday.")
    }
}
```

### Loops

Go has only one looping keyword: `for`. However, it's very versatile and can be used in several ways.

#### The Complete `for` Loop

This is the classic three-component loop found in languages like C or Java.

**Syntax**: `for initialization; condition; post-statement { ... }`

- **initialization**: Executed once before the loop starts (e.g., `i := 0`).
- **condition**: Evaluated before every iteration. The loop stops when it's `false`.
- **post-statement**: Executed at the end of every iteration (e.g., `i++`).

```go
package main

import "fmt"

func main() {
    // Prints numbers from 0 to 4
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

#### The `for` Loop as a "while" Loop

You can omit the initialization and post-statements to make the `for` loop behave like a `while` loop from other languages.

**Syntax**: `for condition { ... }`

```go
package main

import "fmt"

func main() {
    sum := 1
    // This loop continues as long as 'sum' is less than 10
    for sum < 10 {
        sum += sum
        fmt.Println("Current sum:", sum)
    }
    fmt.Println("Final sum:", sum) // Final sum: 16
}
```

#### The Infinite `for` Loop

If you omit the condition entirely, you create an infinite loop. This is common for servers that listen for requests continuously or for loops that have their exit condition handled internally, usually with a `break`.

**Syntax**: `for { ... }`

```go
for {
    fmt.Println("This will run forever...")
    // You would typically have a condition inside to 'break' out of the loop
}
```

### Jump Statements

Jump statements unconditionally transfer control to another part of your code.

#### `break`

The `break` statement immediately **terminates** the innermost `for` loop or `switch` statement.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            fmt.Println("Breaking the loop at 5!")
            break // Exit the loop
        }
        fmt.Println(i)
    }
}
```

**Output:**
```
0
1
2
3
4
Breaking the loop at 5!
```

#### `continue`

The `continue` statement **skips** the rest of the current iteration and proceeds to the **next iteration** of the loop.

```go
package main

import "fmt"

func main() {
    // This loop will only print odd numbers
    for i := 0; i < 10; i++ {
        if i%2 == 0 { // If the number is even...
            continue  // ...skip to the next iteration
        }
        fmt.Println(i)
    }
}
```

**Output:**
```
1
3
5
7
9
```

#### `goto`

The `goto` statement jumps to a predefined label within the same function.

**Warning ⚠️**: The use of `goto` is **heavily discouraged** in modern programming, including Go. It can make code flow confusing, difficult to read, and hard to debug. It is almost always better to use loops and conditional statements.

```go
package main

import "fmt"

func main() {
    i := 0
Here: // This is a label
    fmt.Println(i)
    i++
    if i < 5 {
        goto Here // Jump back to the 'Here' label
    }
    fmt.Println("Done.")
}
```

This example essentially creates a loop that prints numbers from 0 to 4. However, a standard `for` loop would be much clearer and is the recommended approach.