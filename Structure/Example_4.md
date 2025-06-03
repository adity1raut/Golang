package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func main() {
    p := Person{Name: "Alice", Age: 30}
    jsonData , _ := json.Marshal(p)
    fmt.Println(string(jsonData))


    var people []Person
    people = append(people, Person{Name: "Alice", Age: 30})
    people = append(people, Person{Name: "Bob", Age: 25})
    people = append(people, Person{Name: "Aditya", Age: 20})
    fmt.Println(people[2].Name) // "Alice"
    fmt.Println(people)
}