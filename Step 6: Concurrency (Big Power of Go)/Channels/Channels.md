Here is the content formatted for a `README.md` file:

````markdown
# 📘 Go Channels – Basic to Advanced

This guide explains **Go channels** from basic concepts to advanced patterns like fan-in/fan-out.

---

## 🟢 1. What are Channels?

Channels are used to **communicate between goroutines** in Go.

- They allow goroutines to **send and receive** data.
- They help you **synchronize** data without using locks.

```go
ch := make(chan int) // Creating a channel of type int
````

---

## 🟢 2. Sending and Receiving

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello" // send
	}()

	msg := <-ch // receive
	fmt.Println(msg)
}
```

✅ **Output:**

```
Hello
```

---

## 🟢 3. Channel Direction

You can restrict a channel to **only send** or **only receive**.

```go
func sendData(ch chan<- int) {
	ch <- 100 // Only send
}

func receiveData(ch <-chan int) {
	fmt.Println(<-ch) // Only receive
}
```

---

## 🟢 4. Buffered Channels

Buffered channels allow temporary storage of values.

```go
ch := make(chan int, 2) // Buffer size 2
ch <- 1
ch <- 2
fmt.Println(<-ch)
fmt.Println(<-ch)
```

✅ **Output:**

```
1
2
```

---

## 🟢 5. Closing Channels

Use `close(ch)` when no more values will be sent.

```go
ch := make(chan int)
go func() {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}()

for val := range ch {
	fmt.Println(val)
}
```

✅ **Output:**

```
0
1
2
```

---

## 🟢 6. Select Statement

Wait on multiple channel operations using `select`.

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "from ch1" }()
go func() { ch2 <- "from ch2" }()

select {
case msg1 := <-ch1:
	fmt.Println(msg1)
case msg2 := <-ch2:
	fmt.Println(msg2)
}
```

---

## 🟢 7. Channel with WaitGroup

Use channels with `sync.WaitGroup` to wait for goroutines.

```go
var wg sync.WaitGroup
ch := make(chan string)

wg.Add(2)

go func() {
	defer wg.Done()
	ch <- "Hello"
}()

go func() {
	defer wg.Done()
	fmt.Println(<-ch)
}()

wg.Wait()
```

---

## 🟢 8. Fan-Out / Fan-In Patterns

These are concurrency design patterns using channels.

### ✅ Fan-Out

Multiple goroutines read from the same input channel to process work.

### ✅ Fan-In

Multiple goroutines write to a single output channel to merge results.

---

## 📚 Summary

* `chan T` creates a channel of type `T`.
* Use `<-` for sending/receiving data.
* Use `close(ch)` to close channels.
* Use `select` for multiple channel operations.
* Use `sync.WaitGroup` for waiting on goroutines.
* Patterns like Fan-Out and Fan-In are powerful for concurrent design.

```

Would you like me to generate a downloadable file (`README.md`) for you?
```
