package main 

import "fmt"

type Markets interface {
	Shop() string
}

type Chocklet struct {
	name string
}

type Pens struct {
	name string
}

type Packs struct {
	name string
}

func (c Chocklet) Shop() string {
  return "You will orfer Chcocklets"
}

func (p Pens) Shop() string {
	return "You have By a Pens"
}

func (p Packs) Shop() string {
	return "You Byee  a Packs"
}

func main () {
	c := Chocklet{"FFFF"}
	p := Pens{"GGG"}
	j := Packs{"HHH"}



	fmt.Println(c.Shop())
	fmt.Println(p.Shop())
	fmt.Println(j.Shop())
}