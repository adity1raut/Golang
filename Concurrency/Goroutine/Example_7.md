package main

import "fmt"

// This is used for Sending and also reciving data from users can gerrates a mutiples taks at time
func SendOnly(ch chan <- int) {
	ch <- 7852
}

func ReciveOnly(ch <-chan int){
	val := <-ch 
	fmt.Println(val)
}
func main () {
	ch := make(chan string)
	fmt.Println(ch)
}