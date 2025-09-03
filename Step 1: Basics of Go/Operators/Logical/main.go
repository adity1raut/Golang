

package logical

import "fmt"

// Logical operators
func main() {
	a := true
	b := false
	
	fmt.Println("a && b =", a && b)
	fmt.Println("a || b =", a || b)
	fmt.Println("!a =", !a)
	fmt.Println("!b =", !b)
}
