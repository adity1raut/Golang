

package main

func main(){

	ch := make(chan string)  //creating a channel of type string you can used any typr here

	go func ()  {
		ch <- "Hello from Go Channel!" // <- Used This Oprators for sending a value to channels 
	}()

	msg := <- ch // <- Operator also used for receiving a value from channel
	println(msg)
}