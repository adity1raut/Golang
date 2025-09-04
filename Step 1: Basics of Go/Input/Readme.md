## 🔹 1. Using `fmt.Scan` / `fmt.Scanln` / `fmt.Scanf`
The **simplest way** to take input.

### Example: Single input
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name) // input stored in name
    fmt.Println("Hello,", name)
}
```

### Example: Multiple inputs
```go
var a, b int
fmt.Print("Enter two numbers: ")
fmt.Scan(&a, &b)
fmt.Println("Sum:", a+b)
```

### Difference:
- `Scan` → stops at space/newline.  
- `Scanln` → reads until newline.  
- `Scanf` → formatted input like C’s `scanf`.

```go
var name string
var age int
fmt.Scanf("%s %d", &name, &age)
fmt.Println("Name:", name, "Age:", age)
```

---

## 🔹 2. Using `bufio.Reader` (for full line input)
Best for **reading a complete line with spaces**.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a sentence: ")
    input, _ := reader.ReadString('\n')
    fmt.Println("You entered:", input)
}
```

---

## 🔹 3. Using `os.Stdin` with `fmt.Fscan`
Good when you want to read structured input.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    var x, y int
    fmt.Print("Enter two numbers: ")
    fmt.Fscan(os.Stdin, &x, &y)
    fmt.Println("Product:", x*y)
}
```

---

## 🔹 4. Using `bufio.Scanner` (for token-based input)
Efficient for competitive programming / reading multiple inputs.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter text: ")
    scanner.Scan()
    text := scanner.Text()
    fmt.Println("Input:", text)
}
```

---

## 🔹 Input Conversion
All inputs come as **strings**, so often you need to convert them.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a number: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)   // remove newline
    num, _ := strconv.Atoi(input)      // string → int
    fmt.Println("Square:", num*num)
}
```

---

## 🔑 Key Points
1. `fmt.Scan` → quick, but stops at space/newline.  
2. `fmt.Scanln` → reads until newline.  
3. `fmt.Scanf` → formatted input.  
4. `bufio.Reader` → reads line (with spaces).  
5. `bufio.Scanner` → best for tokenized input.  
6. Use `strconv` to convert strings to `int`, `float64`, etc.  

