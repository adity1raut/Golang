

package main 

import "fmt"

func main() {

	m := make(map[string]int) 

	m["Aditya Raut "] = 21
	m["Montya Raut "] = 18 

	m["Raja Raut "] = 16 

	fmt.Println(m)

	m["Shinggham"] = 98 
	m["Modi"] = 10 

	fmt.Println(m)

	delete(m , "Modi")

	fmt.Println(m)
}