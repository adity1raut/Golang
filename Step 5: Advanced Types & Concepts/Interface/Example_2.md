package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	var d Dog
	var c Cat

	makeSound(d)
	makeSound(c)
}
