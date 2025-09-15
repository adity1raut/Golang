```go

package main

import "fmt"

func main() {
	//Maps in a Golangs details explinantions With and Examples

	// String to int map
	ages := make(map[string]int)
	fmt.Println(ages)

	// Literal initialization
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)

	// Zero value is nil
	var nilMap map[string]int // nil map (cannot be used directly)
	fmt.Print(nilMap)
}
