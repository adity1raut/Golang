//Interface in a golang Basic

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base float64
	height float64
}

type Reactangle struct {
	length float64
	breath float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Reactangle) Area() float64 {
	return r.length * r.breath
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func PrintArea (s Shape ) {
	fmt.Println("Area is :" , s.Area())
}

func main () {
	c := Circle{10}
	r := Reactangle{10 , 20 }
	t := Triangle{12 , 00}

	PrintArea(c)
	PrintArea(r)
	PrintArea(t)


}