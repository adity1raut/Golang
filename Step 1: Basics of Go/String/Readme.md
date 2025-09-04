

## 🔹 What is a String in Go?

* In Go, a **string is a sequence of bytes** (not characters).
* Strings are **immutable** → once created, you cannot change them.
* Internally, a string is stored as a **read-only slice of bytes**.

```go
package main

import "fmt"

func main() {
    str := "Hello, Go!"
    fmt.Println(str)
}
```

---

## 🔹 How Strings are Stored

* Go stores strings as **UTF-8 encoded bytes**.
* Each character (rune) may take **1 to 4 bytes**.
* For example:

  * English letters → 1 byte
  * Hindi/Chinese characters → more than 1 byte

```go
package main

import "fmt"

func main() {
    str := "Hello"
    fmt.Println(len(str)) // 5 (each character is 1 byte)

    str2 := "हाय"
    fmt.Println(len(str2)) // 6 (each Hindi character takes 3 bytes)
}
```

---

## 🔹 String vs Rune

* **String** → sequence of bytes.
* **Rune** → represents a Unicode code point (int32).

```go
package main

import "fmt"

func main() {
    word := "हाय"
    for i, r := range word {
        fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
    }
}
```

Output:

```
Index: 0, Rune: ह, Unicode: U+0939
Index: 3, Rune: ा, Unicode: U+093E
Index: 6, Rune: य, Unicode: U+092F
```

👉 Notice: index jumps by 3 because each character takes 3 bytes.

---

## 🔹 Common String Operations

### 1. Concatenation

```go
a := "Hello"
b := "World"
c := a + " " + b
fmt.Println(c) // Hello World
```

### 2. Indexing (returns byte, not character)

```go
s := "Go"
fmt.Println(s[0])      // 71 (ASCII value of 'G')
fmt.Println(string(s[0])) // "G"
```

### 3. Slicing

```go
text := "Hello, Go!"
fmt.Println(text[0:5]) // Hello
```

### 4. Iterating

* By **bytes**:

```go
for i := 0; i < len(text); i++ {
    fmt.Printf("%c ", text[i])
}
```

* By **runes (characters)**:

```go
for _, r := range text {
    fmt.Printf("%c ", r)
}
```

---

## 🔹 Useful String Functions (from `strings` package)

```go
import "strings"

// Contains
strings.Contains("hello world", "world") // true

// ToUpper, ToLower
strings.ToUpper("go") // "GO"

// Split
strings.Split("a,b,c", ",") // ["a", "b", "c"]

// Replace
strings.Replace("foo bar foo", "foo", "baz", -1) // "baz bar baz"

// Trim
strings.Trim("  hello  ", " ") // "hello"

// HasPrefix / HasSuffix
strings.HasPrefix("golang", "go") // true
strings.HasSuffix("golang", "lang") // true
```

---

## 🔹 String Conversion

Go often converts between **string**, **\[]byte**, and **\[]rune**.

```go
str := "GoLang"

// string → []byte
b := []byte(str)
fmt.Println(b) // [71 111 76 97 110 103]

// string → []rune
r := []rune(str)
fmt.Println(r) // [71 111 76 97 110 103]

// []byte → string
s := string(b)
fmt.Println(s) // GoLang
```

---

## 🔹 Key Points

1. Strings are **immutable** in Go.
2. Strings are **UTF-8 encoded** byte sequences.
3. Use **runes** for Unicode characters.
4. `strings` package provides powerful functions.
5. Convert between `string`, `[]byte`, and `[]rune` when needed.

---

