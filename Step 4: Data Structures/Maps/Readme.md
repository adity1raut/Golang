
## **1. What is a map in Go?**

A **map** is a built-in **key-value data structure** in Go. Think of it like a dictionary in Python or HashMap in Java.

* **Key**: Unique identifier
* **Value**: Data associated with the key

**Syntax:**

```go
var myMap map[keyType]valueType
```

**Example:**

```go
var scores map[string]int
```

Here, `scores` maps a `string` (like a name) to an `int` (like marks).

---

## **2. Creating a map**

You cannot directly use a map without initializing it. Use `make`:

```go
scores := make(map[string]int)
```

Or using a **map literal**:

```go
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
}
```

---

## **3. Adding or updating elements**

```go
scores["Charlie"] = 95   // Add new key-value
scores["Alice"] = 99     // Update existing key
```

---

## **4. Accessing elements**

```go
fmt.Println(scores["Alice"]) // Output: 99
```

If the key does **not exist**, Go returns the **zero value** of the type (`0` for int, `""` for string, `false` for bool).

You can also check if a key exists:

```go
value, ok := scores["Daniel"]
if ok {
    fmt.Println("Found:", value)
} else {
    fmt.Println("Not found")
}
```

* `ok` is `true` if key exists, otherwise `false`.

---

## **5. Deleting elements**

```go
delete(scores, "Bob")
```

`delete()` removes a key-value pair. If the key doesn’t exist, it does nothing (no error).

---

## **6. Iterating over a map**

You can loop through all key-value pairs using `range`:

```go
for name, score := range scores {
    fmt.Println(name, score)
}
```

If you only need keys:

```go
for name := range scores {
    fmt.Println(name)
}
```

---

## **7. Nested maps**

Maps can contain other maps:

```go
students := map[string]map[string]int{
    "Alice": {"Math": 90, "Science": 95},
    "Bob":   {"Math": 85, "Science": 80},
}

fmt.Println(students["Alice"]["Science"]) // 95
```

---

## **8. Map of slices**

You can also store slices as values:

```go
emails := map[string][]string{
    "Alice": {"alice@gmail.com", "alice@yahoo.com"},
}

emails["Bob"] = []string{"bob@gmail.com"}
```

---

## **9. Passing maps to functions**

Maps are **reference types**, so changes inside a function affect the original map.

```go
func updateScore(scores map[string]int, name string, score int) {
    scores[name] = score
}

updateScore(scores, "Alice", 100)
fmt.Println(scores["Alice"]) // 100
```

---

## **10. Advanced: Map performance & concurrency**

* Maps are **fast** for lookups (`O(1)` average)
* **Zero value of a map is `nil`**. Using a nil map for writing will panic.
* Maps are **not safe for concurrent writes**. Use `sync.Map` for thread-safe operations.

**Example of sync.Map:**

```go
import "sync"

var m sync.Map
m.Store("Alice", 100)
value, ok := m.Load("Alice")
fmt.Println(value, ok)
```

---

## **11. Advanced: Sorting a map by keys or values**

Maps are **unordered** in Go. To sort:

```go
import "sort"

keys := make([]string, 0, len(scores))
for k := range scores {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, scores[k])
}
```

---

## ✅ Summary

* Map = key-value pair (`map[keyType]valueType`)
* Must be **initialized** using `make` or literal
* Accessing a key returns zero value if key does not exist
* `delete()` removes a key
* Iteration is done with `range`
* Maps are **reference types**, not safe for concurrent writes
* Can use `sync.Map` for concurrency

---