package main 

import "fmt"

type Student struct {
    name string 
    age int 
    isPass bool
}

func Congratulations (s Student){
    fmt.Println(s.name , " Is Pass ")
}

func (s Student) String() string {
    return fmt.Sprintf("Name: %s, Age: %d, IsPass: %t", s.name, s.age, s.isPass)
}

func main () {
    s1 := Student {"Aditya Raut" , 20, true}
    s2 := Student {"Raj Raut" , 17, true}

    Congratulations(s1)

    fmt.Println(s2)
   

}