package main

import "fmt"

type student struct {
	student  string
	age      int
	isPassed bool
}

func Passing (name string , age int) *student {
	return &student{
		student:  name,
		age:      age,
		isPassed: true,
	}
}

func main() {
	s := student{"Aditya Raut", 20, true}
	fmt.Println(s.student) // Output: Aditya Raut
	fmt.Println(s.age)      // Output: 20
	fmt.Println(s.isPassed) // Output: true
	fmt.Println(s) // Output: {Aditya Raut 20 true}

	s2 := Passing("Aditya Raut", 20)
	fmt.Println(s2.student) // Output: Aditya Raut
	fmt.Println(s2.age)      // Output: 20
	fmt.Println(s2.isPassed) // Output: true
	fmt.Println(s2) // Output: &{Aditya Raut 20 true} // Pointer to struct Becouse Vo Direactllly Address asinge karega Variable or Function ka!



}