package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	// 1. Creating and initializing maps
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Println("Initial scores:", scores)

	// 2. Adding or updating elements
	scores["Charlie"] = 95
	scores["Alice"] = 99
	fmt.Println("After adding/updating:", scores)

	// 3. Accessing elements
	fmt.Println("Alice's score:", scores["Alice"])

	// 4. Checking if key exists
	if val, ok := scores["Daniel"]; ok {
		fmt.Println("Daniel's score:", val)
	} else {
		fmt.Println("Daniel not found")
	}

	// 5. Deleting elements
	delete(scores, "Bob")
	fmt.Println("After deleting Bob:", scores)

	// 6. Iterating over map
	fmt.Println("Iterating over map:")
	for name, score := range scores {
		fmt.Println(name, score)
	}

	// 7. Nested maps
	students := map[string]map[string]int{
		"Alice": {"Math": 90, "Science": 95},
		"Bob":   {"Math": 85, "Science": 80},
	}
	fmt.Println("Alice's Science score:", students["Alice"]["Science"])

	// 8. Map of slices
	emails := map[string][]string{
		"Alice": {"alice@gmail.com", "alice@yahoo.com"},
	}
	emails["Bob"] = []string{"bob@gmail.com"}
	fmt.Println("Emails map:", emails)

	// 9. Passing map to function
	updateScore(scores, "Alice", 100)
	fmt.Println("After function update:", scores)

	// 10. Advanced: Sorting map by keys
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Sorted scores by name:")
	for _, k := range keys {
		fmt.Println(k, scores[k])
	}

	// 11. Advanced: Concurrency safe map using sync.Map
	var m sync.Map
	m.Store("Alice", 100)
	m.Store("Bob", 80)
	value, ok := m.Load("Alice")
	if ok {
		fmt.Println("sync.Map Alice:", value)
	}
	m.Delete("Bob")
	m.Range(func(k, v any) bool {
		fmt.Println("sync.Map key:", k, "value:", v)
		return true
	})
}

func updateScore(scores map[string]int, name string, score int) {
	scores[name] = score
}
