package main 

import "fmt"

func main(){
   var num int
   fmt.Println("Enter a number:")
   fmt.Scan(&num)

   switch num {
   case 1: 
      fmt.Print("One")
   case 2:
      fmt.Print("Two")
   case 3:
	  fmt.Print("Three")
   default:
      fmt.Print("Number not recognized")
   }
}